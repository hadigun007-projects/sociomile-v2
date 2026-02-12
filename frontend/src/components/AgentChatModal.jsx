import { useState, useRef, useEffect } from 'react';
import { useMutation, useQuery, useQueryClient } from '@tanstack/react-query';
import api from '../api/axios';
import { Card } from './Card';
import { Button } from './Button';

const AgentChatModal = ({ isOpen, onClose, conversation: initialConversation }) => {
    const [message, setMessage] = useState('');
    const messagesEndRef = useRef(null);
    const queryClient = useQueryClient();

    // Fetch live conversation data
    const { data: conversation } = useQuery({
        queryKey: ['conversation', initialConversation?.id],
        queryFn: async () => {
            const response = await api.get(`/conversations/${initialConversation.id}`);
            return response.data.data;
        },
        enabled: isOpen && !!initialConversation?.id,
        refetchInterval: 3000, // Auto-refresh every 3 seconds to get new messages
    });

    const messages = (conversation?.messages || []).sort((a, b) =>
        new Date(a.created_at) - new Date(b.created_at)
    );

    const scrollToBottom = () => {
        messagesEndRef.current?.scrollIntoView({ behavior: 'smooth' });
    };

    useEffect(() => {
        if (isOpen && messages.length > 0) {
            setTimeout(scrollToBottom, 100);
        }
    }, [isOpen, messages.length]);

    const replyMutation = useMutation({
        mutationFn: async ({ conversationId, message }) => {
            const response = await api.post(`/conversations/${conversationId}/reply`, {
                message,
            });
            return response.data;
        },
        onSuccess: () => {
            setMessage('');
            // Invalidate both queries to refresh data
            queryClient.invalidateQueries(['conversation', initialConversation?.id]);
            queryClient.invalidateQueries(['conversations']);
        },
        onError: (err) => {
            alert(`Failed to send message: ${err.response?.data?.error || err.message}`);
        },
    });

    const handleSubmit = (e) => {
        e.preventDefault();
        if (!message.trim()) return;

        replyMutation.mutate({
            conversationId: initialConversation.id,
            message: message.trim(),
        });
    };

    if (!isOpen || !initialConversation) return null;

    // Use conversation from query if available, otherwise use initialConversation
    const displayConversation = conversation || initialConversation;

    return (
        <div className="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center z-50 p-4">
            <Card className="w-full max-w-3xl h-[80vh] flex flex-col">
                {/* Header */}
                <div className="flex items-center justify-between mb-4 pb-4 border-b">
                    <div>
                        <h2 className="text-2xl font-bold text-gray-800">Conversation</h2>
                        <div className="flex gap-4 mt-2 text-sm text-gray-600">
                            <span className="font-mono">ID: {displayConversation.id.slice(0, 8)}...</span>
                            <span className={`px-2 py-0.5 rounded-full text-xs font-medium ${displayConversation.status === 'open' ? 'bg-green-100 text-green-700' :
                                displayConversation.status === 'assigned' ? 'bg-blue-100 text-blue-700' :
                                    'bg-gray-100 text-gray-700'
                                }`}>
                                {displayConversation.status}
                            </span>
                        </div>
                    </div>
                    <button
                        onClick={onClose}
                        className="text-gray-400 hover:text-gray-600 text-3xl leading-none"
                    >
                        ×
                    </button>
                </div>

                {/* Messages Container */}
                <div className="flex-1 overflow-y-auto mb-4 space-y-3 px-2">
                    {messages.length === 0 ? (
                        <div className="text-center text-gray-400 py-8">
                            No messages in this conversation
                        </div>
                    ) : (
                        messages.map((msg) => (
                            <div
                                key={msg.id}
                                className={`flex ${msg.sender_type === 'customer' ? 'justify-start' : 'justify-end'}`}
                            >
                                <div
                                    className={`max-w-[70%] rounded-lg px-4 py-2 ${msg.sender_type === 'customer'
                                        ? 'bg-gray-100 text-gray-800'
                                        : 'bg-purple-600 text-white'
                                        }`}
                                >
                                    <div className="text-xs opacity-70 mb-1">
                                        {msg.sender_type === 'customer' ? 'Customer' : 'You'}
                                    </div>
                                    <div className="text-sm whitespace-pre-wrap break-words">
                                        {msg.message}
                                    </div>
                                    <div className="text-xs opacity-60 mt-1">
                                        {new Date(msg.created_at).toLocaleString('id-ID')}
                                    </div>
                                </div>
                            </div>
                        ))
                    )}
                    <div ref={messagesEndRef} />
                </div>

                {/* Message Input */}
                <form onSubmit={handleSubmit} className="flex gap-2 pt-4 border-t">
                    <input
                        type="text"
                        value={message}
                        onChange={(e) => setMessage(e.target.value)}
                        placeholder="Type your message..."
                        className="flex-1 px-4 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-purple-500"
                        disabled={replyMutation.isPending}
                    />
                    <Button
                        type="submit"
                        isLoading={replyMutation.isPending}
                        disabled={!message.trim() || replyMutation.isPending}
                        className="px-6"
                    >
                        Send
                    </Button>
                </form>
            </Card>
        </div>
    );
};

export default AgentChatModal;
