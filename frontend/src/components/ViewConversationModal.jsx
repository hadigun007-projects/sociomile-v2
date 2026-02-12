import { Card } from './Card';

const ViewConversationModal = ({ isOpen, onClose, conversation }) => {
    if (!isOpen || !conversation) return null;

    const messages = conversation.messages || [];

    return (
        <div className="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center z-50 p-4">
            <Card className="w-full max-w-3xl max-h-[80vh] flex flex-col">
                {/* Header */}
                <div className="flex items-center justify-between mb-4 pb-4 border-b">
                    <div>
                        <h2 className="text-2xl font-bold text-gray-800">View Conversation</h2>
                        <div className="flex gap-4 mt-2 text-sm text-gray-600">
                            <span className="font-mono">ID: {conversation.id.slice(0, 8)}...</span>
                            <span className={`px-2 py-0.5 rounded-full text-xs font-medium ${conversation.status === 'open' ? 'bg-green-100 text-green-700' :
                                    conversation.status === 'assigned' ? 'bg-blue-100 text-blue-700' :
                                        'bg-gray-100 text-gray-700'
                                }`}>
                                {conversation.status}
                            </span>
                            {conversation.assigned_agent && (
                                <span className="text-purple-600">
                                    Agent: {conversation.assigned_agent.email}
                                </span>
                            )}
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
                                        {msg.sender_type === 'customer' ? 'Customer' : 'Agent'}
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
                </div>

                {/* Info Footer */}
                <div className="pt-4 border-t bg-yellow-50 -mx-6 -mb-6 px-6 py-4 rounded-b-lg">
                    <div className="flex items-center gap-2 text-sm text-yellow-800">
                        <svg className="w-5 h-5" fill="currentColor" viewBox="0 0 20 20">
                            <path fillRule="evenodd" d="M18 10a8 8 0 11-16 0 8 8 0 0116 0zm-7-4a1 1 0 11-2 0 1 1 0 012 0zM9 9a1 1 0 000 2v3a1 1 0 001 1h1a1 1 0 100-2v-3a1 1 0 00-1-1H9z" clipRule="evenodd" />
                        </svg>
                        <span>Read-only mode: You can view this conversation but cannot send messages.</span>
                    </div>
                </div>
            </Card>
        </div>
    );
};

export default ViewConversationModal;
