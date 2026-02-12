import { useState, useEffect, useRef } from 'react';
import { useMutation, useQuery } from '@tanstack/react-query';
import api from '../api/axios';
import { Card } from '../components/Card';
import Input from '../components/Input';
import { Button } from '../components/Button';

const ChannelSimulator = () => {
    const [mode, setMode] = useState('new'); // 'new' or 'existing'
    const [tenantId, setTenantId] = useState('');
    const [customerExternalId, setCustomerExternalId] = useState('');
    const [conversationId, setConversationId] = useState('');
    const [message, setMessage] = useState('');
    const [responseMessage, setResponseMessage] = useState('');
    const [loadedConversation, setLoadedConversation] = useState(null);
    const messagesEndRef = useRef(null);

    // Auto scroll to bottom when messages change
    const scrollToBottom = () => {
        messagesEndRef.current?.scrollIntoView({ behavior: "smooth" });
    };

    useEffect(() => {
        scrollToBottom();
    }, [loadedConversation?.messages]);

    // Fetch tenants for dropdown
    const { data: tenants = [], isLoading: tenantsLoading } = useQuery({
        queryKey: ['public-tenants'],
        queryFn: async () => {
            const response = await api.get('/tenants/public');
            return response.data.data;
        },
    });

    const webhookMutation = useMutation({
        mutationFn: async (data) => {
            const response = await api.post('/channel/webhook', data);
            return response.data;
        },
        onSuccess: (data, variables) => {
            const convId = variables.conversation_id || data.data.conversation_id;
            if (convId) {
                // If sending to existing conversation or just created one, reload it
                setResponseMessage('Message sent successfully!');
                setMessage('');
                loadConversationMutation.mutate(convId);
                if (!conversationId) setConversationId(convId);
                if (mode === 'new') setMode('existing');
            } else {
                setResponseMessage(`Message sent successfully! Customer ID: ${data.data.customer_external_id}`);
                setMessage('');
            }
        },
        onError: (err) => {
            setResponseMessage(`Error: ${err.response?.data?.error || err.message}`);
        },
    });

    const loadConversationMutation = useMutation({
        mutationFn: async (convId) => {
            const response = await api.get(`/conversations/${convId}/public`);
            return response.data.data;
        },
        onSuccess: (data) => {
            setLoadedConversation(data);
            setResponseMessage('Conversation loaded successfully!');
        },
        onError: (err) => {
            setResponseMessage(`Error: ${err.response?.data?.error || err.message}`);
            setLoadedConversation(null);
        },
    });

    const handleSubmit = (e) => {
        e.preventDefault();
        setResponseMessage('');

        if (mode === 'new') {
            webhookMutation.mutate({
                tenant_id: tenantId,
                customer_external_id: customerExternalId,
                message: message,
            });
        }
    };

    const handleLoadConversation = (e) => {
        e.preventDefault();
        setResponseMessage('');
        setLoadedConversation(null);
        loadConversationMutation.mutate(conversationId);
    };

    const handleSendToConversation = (e) => {
        e.preventDefault();
        if (!message.trim() || !loadedConversation) return;

        setResponseMessage('');
        webhookMutation.mutate({
            tenant_id: loadedConversation.tenant_id, // Use ID from loaded data
            conversation_id: loadedConversation.id,
            message: message.trim(),
        });
    };

    return (
        <div className="min-h-screen bg-gradient-to-br from-purple-50 via-white to-purple-50 p-6">
            <div className="max-w-6xl mx-auto">
                <div className="mb-8 text-center">
                    <h1 className="text-4xl font-bold text-gray-800 mb-2">Channel Webhook Simulator</h1>
                    <p className="text-gray-600">
                        Simulate incoming messages from external channels (WhatsApp, Instagram, etc.)
                    </p>
                </div>

                <div className="grid grid-cols-1 lg:grid-cols-2 gap-6">
                    {/* Left Card - Controls */}
                    <Card className="bg-white">
                        <h2 className="text-xl font-semibold text-gray-800 mb-4">Controls</h2>

                        {/* Mode Selector */}
                        <div className="mb-4 flex gap-2">
                            <button
                                type="button"
                                onClick={() => {
                                    setMode('new');
                                    setLoadedConversation(null);
                                }}
                                className={`flex-1 px-4 py-2 rounded-lg font-medium transition-colors ${mode === 'new'
                                    ? 'bg-purple-600 text-white'
                                    : 'bg-gray-100 text-gray-600 hover:bg-gray-200'
                                    }`}
                            >
                                New Message
                            </button>
                            <button
                                type="button"
                                onClick={() => {
                                    setMode('existing');
                                    setLoadedConversation(null);
                                }}
                                className={`flex-1 px-4 py-2 rounded-lg font-medium transition-colors ${mode === 'existing'
                                    ? 'bg-purple-600 text-white'
                                    : 'bg-gray-100 text-gray-600 hover:bg-gray-200'
                                    }`}
                            >
                                Load Conversation
                            </button>
                        </div>

                        {mode === 'new' ? (
                            <form onSubmit={handleSubmit} className="space-y-4">
                                <div>
                                    <label className="block text-sm font-medium text-gray-700 mb-1">
                                        Tenant
                                    </label>
                                    <select
                                        value={tenantId}
                                        onChange={(e) => setTenantId(e.target.value)}
                                        required
                                        disabled={tenantsLoading}
                                        className="block w-full rounded-lg border border-gray-300 px-3 py-2 shadow-sm focus:border-purple-500 focus:outline-none focus:ring-1 focus:ring-purple-500"
                                    >
                                        <option value="">Select a tenant...</option>
                                        {tenants.map((tenant) => (
                                            <option key={tenant.id} value={tenant.id}>
                                                {tenant.name}
                                            </option>
                                        ))}
                                    </select>
                                </div>

                                <Input
                                    label="Customer External ID"
                                    type="text"
                                    value={customerExternalId}
                                    onChange={(e) => setCustomerExternalId(e.target.value)}
                                    required
                                    placeholder="e.g., +6281234567890"
                                />

                                <div>
                                    <label className="block text-sm font-medium text-gray-700 mb-1">
                                        Message
                                    </label>
                                    <textarea
                                        value={message}
                                        onChange={(e) => setMessage(e.target.value)}
                                        required
                                        rows={4}
                                        className="block w-full rounded-lg border border-gray-300 px-3 py-2 shadow-sm focus:border-purple-500 focus:outline-none focus:ring-1 focus:ring-purple-500"
                                        placeholder="Enter customer message..."
                                    />
                                </div>

                                <Button
                                    type="submit"
                                    className="w-full"
                                    isLoading={webhookMutation.isPending}
                                >
                                    Send Message
                                </Button>
                            </form>
                        ) : (
                            <form onSubmit={handleLoadConversation} className="space-y-4">
                                <Input
                                    label="Conversation ID"
                                    type="text"
                                    value={conversationId}
                                    onChange={(e) => setConversationId(e.target.value)}
                                    required
                                    placeholder="Enter conversation ID to view"
                                />

                                <Button
                                    type="submit"
                                    className="w-full"
                                    isLoading={loadConversationMutation.isPending}
                                >
                                    Load Conversation
                                </Button>
                            </form>
                        )}

                        {responseMessage && (
                            <div
                                className={`mt-4 p-3 rounded-lg text-sm font-medium ${responseMessage.includes('Error')
                                    ? 'bg-red-50 text-red-700'
                                    : 'bg-green-50 text-green-700'
                                    }`}
                            >
                                {responseMessage}
                            </div>
                        )}
                    </Card>

                    {/* Right Card - Chat Interface or Instructions */}
                    {loadedConversation ? (
                        <Card className="bg-white flex flex-col h-[600px]">
                            {/* Chat Header */}
                            <div className="border-b pb-3 mb-4">
                                <div className="flex items-center justify-between">
                                    <h2 className="text-lg font-semibold text-gray-800">Chat</h2>
                                    <span className={`text-xs font-medium capitalize px-2 py-1 rounded ${loadedConversation.status === 'open' ? 'bg-green-100 text-green-700' :
                                        loadedConversation.status === 'assigned' ? 'bg-blue-100 text-blue-700' :
                                            'bg-gray-100 text-gray-700'
                                        }`}>
                                        {loadedConversation.status}
                                    </span>
                                </div>
                                <p className="text-xs text-gray-500 mt-1 font-mono">{conversationId}</p>
                            </div>

                            {/* Messages Area */}
                            <div className="flex-1 overflow-y-auto space-y-3 mb-4 px-2">
                                {/* Ticket Info */}
                                {loadedConversation.ticket && (
                                    <div className="mb-4 bg-purple-50 p-3 rounded-lg border border-purple-100 flex items-center justify-between">
                                        <div>
                                            <div className="text-xs text-purple-600 font-bold uppercase tracking-wide">
                                                Ticket #{loadedConversation.ticket.id.slice(0, 8)}
                                            </div>
                                            <div className="text-sm font-medium text-gray-800">
                                                {loadedConversation.ticket.title}
                                            </div>
                                        </div>
                                        <div className={`px-2 py-1 rounded text-xs font-bold uppercase ${loadedConversation.ticket.status === 'resolved' || loadedConversation.ticket.status === 'closed'
                                            ? 'bg-green-100 text-green-700'
                                            : 'bg-yellow-100 text-yellow-700'
                                            }`}>
                                            {loadedConversation.ticket.status.replace('_', ' ')}
                                        </div>
                                    </div>
                                )}

                                {loadedConversation.messages?.length === 0 ? (
                                    <div className="text-center text-gray-500 py-8">
                                        No messages yet
                                    </div>
                                ) : (
                                    [...(loadedConversation.messages || [])]
                                        .sort((a, b) => new Date(a.created_at) - new Date(b.created_at))
                                        .map((msg, idx) => (
                                            <div
                                                key={msg.id || idx}
                                                className={`flex ${msg.sender_type === 'customer' ? 'justify-end' : 'justify-start'}`}
                                            >
                                                <div
                                                    className={`max-w-[70%] rounded-lg px-4 py-2 ${msg.sender_type === 'customer'
                                                        ? 'bg-purple-600 text-white'
                                                        : 'bg-gray-100 text-gray-800'
                                                        }`}
                                                >
                                                    <div className={`text-xs mb-1 font-medium ${msg.sender_type === 'customer' ? 'text-purple-200' : 'text-gray-500'
                                                        }`}>
                                                        {msg.sender_type === 'customer' ? 'Customer' : 'Agent'}
                                                    </div>
                                                    <div className="text-sm break-words">{msg.message}</div>
                                                    {msg.created_at && (
                                                        <div className={`text-xs mt-1 ${msg.sender_type === 'customer' ? 'text-purple-200' : 'text-gray-400'
                                                            }`}>
                                                            {new Date(msg.created_at).toLocaleTimeString('id-ID', {
                                                                hour: '2-digit',
                                                                minute: '2-digit'
                                                            })}
                                                        </div>
                                                    )}
                                                </div>
                                            </div>
                                        ))
                                )}
                                <div ref={messagesEndRef} />
                            </div>

                            {/* Message Input */}
                            <div className="border-t pt-4">
                                {loadedConversation.ticket?.status === 'resolved' || loadedConversation.ticket?.status === 'closed' ? (
                                    <div className="text-center py-2 bg-gray-50 rounded-lg text-gray-500 text-sm italic border border-gray-200">
                                        This conversation has been marked as resolved. You cannot send further messages.
                                    </div>
                                ) : (
                                    <form onSubmit={handleSendToConversation} className="flex gap-2">
                                        <input
                                            type="text"
                                            value={message}
                                            onChange={(e) => setMessage(e.target.value)}
                                            required
                                            className="flex-1 rounded-lg border border-gray-300 px-4 py-2 focus:border-purple-500 focus:outline-none focus:ring-1 focus:ring-purple-500"
                                            placeholder="Type a message..."
                                        />
                                        <Button
                                            type="submit"
                                            isLoading={webhookMutation.isPending}
                                            className="px-6"
                                        >
                                            Send
                                        </Button>
                                    </form>
                                )}
                            </div>
                        </Card>
                    ) : (
                        <Card className="bg-purple-50">
                            <h2 className="text-xl font-semibold text-gray-800 mb-4">How It Works</h2>
                            <div className="space-y-3 text-sm text-gray-700">
                                <div className="flex items-start">
                                    <span className="inline-flex items-center justify-center h-6 w-6 rounded-full bg-purple-600 text-white font-bold mr-3 flex-shrink-0">
                                        1
                                    </span>
                                    <p>
                                        <strong>New Message:</strong> Enter customer external ID to create/find customer and conversation, then send a message
                                    </p>
                                </div>
                                <div className="flex items-start">
                                    <span className="inline-flex items-center justify-center h-6 w-6 rounded-full bg-purple-600 text-white font-bold mr-3 flex-shrink-0">
                                        2
                                    </span>
                                    <p>
                                        <strong>Load Conversation:</strong> Enter conversation ID to view the chat interface on the right
                                    </p>
                                </div>
                                <div className="flex items-start">
                                    <span className="inline-flex items-center justify-center h-6 w-6 rounded-full bg-purple-600 text-white font-bold mr-3 flex-shrink-0">
                                        3
                                    </span>
                                    <p>
                                        <strong>Chat Interface:</strong> Your messages (customer) appear on the right (purple), agent messages on the left (gray)
                                    </p>
                                </div>
                            </div>

                            <div className="mt-6 p-3 bg-white rounded-lg border border-purple-200">
                                <p className="text-xs text-gray-600">
                                    <strong>Note:</strong> Get conversation IDs from the database after creating new messages, then load them to see the WhatsApp-style chat interface.
                                </p>
                            </div>
                        </Card>
                    )}
                </div>
            </div>
        </div>
    );
};

export default ChannelSimulator;
