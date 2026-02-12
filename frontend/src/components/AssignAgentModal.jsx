import { useState, useEffect } from 'react';
import { useQuery, useMutation, useQueryClient } from '@tanstack/react-query';
import api from '../api/axios';
import { Card } from './Card';
import { Button } from './Button';

const AssignAgentModal = ({ isOpen, onClose, conversation }) => {
    const [selectedAgentId, setSelectedAgentId] = useState('');
    const queryClient = useQueryClient();

    // Fetch agents from the same tenant
    const { data: users = [], isLoading: usersLoading } = useQuery({
        queryKey: ['users'],
        queryFn: async () => {
            const response = await api.get('/users');
            return response.data.data;
        },
        enabled: isOpen,
    });

    // Filter only agents
    const agents = users.filter(user => user.role === 'agent');

    const assignMutation = useMutation({
        mutationFn: async ({ conversationId, agentId }) => {
            const response = await api.put(`/conversations/${conversationId}/assign`, {
                agent_id: agentId,
            });
            return response.data;
        },
        onSuccess: () => {
            queryClient.invalidateQueries(['conversations']);
            onClose();
        },
        onError: (err) => {
            alert(`Failed to assign agent: ${err.response?.data?.error || err.message}`);
        },
    });

    const handleSubmit = (e) => {
        e.preventDefault();
        if (!selectedAgentId) {
            alert('Please select an agent');
            return;
        }
        assignMutation.mutate({
            conversationId: conversation.id,
            agentId: selectedAgentId,
        });
    };

    useEffect(() => {
        if (isOpen) {
            setSelectedAgentId('');
        }
    }, [isOpen]);

    if (!isOpen) return null;

    return (
        <div className="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center z-50">
            <Card className="w-full max-w-md mx-4">
                <div className="flex items-center justify-between mb-6">
                    <h2 className="text-2xl font-bold text-gray-800">Assign Agent</h2>
                    <button
                        onClick={onClose}
                        className="text-gray-400 hover:text-gray-600 text-2xl leading-none"
                    >
                        ×
                    </button>
                </div>

                <div className="mb-4 p-3 bg-purple-50 rounded-lg">
                    <div className="text-sm text-gray-600">Conversation ID</div>
                    <div className="font-mono text-xs text-purple-700 mt-1">
                        {conversation?.id}
                    </div>
                </div>

                <form onSubmit={handleSubmit} className="space-y-4">
                    <div>
                        <label className="block text-sm font-medium text-gray-700 mb-2">
                            Select Agent
                        </label>
                        {usersLoading ? (
                            <div className="text-sm text-gray-500">Loading agents...</div>
                        ) : agents.length === 0 ? (
                            <div className="text-sm text-gray-500">No agents available in this tenant</div>
                        ) : (
                            <select
                                value={selectedAgentId}
                                onChange={(e) => setSelectedAgentId(e.target.value)}
                                required
                                className="w-full px-3 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-purple-500"
                            >
                                <option value="">Choose an agent...</option>
                                {agents.map((agent) => (
                                    <option key={agent.id} value={agent.id}>
                                        {agent.email}
                                    </option>
                                ))}
                            </select>
                        )}
                    </div>

                    <div className="flex gap-3">
                        <Button
                            type="button"
                            onClick={onClose}
                            className="flex-1 bg-gray-200 text-gray-700 hover:bg-gray-300"
                        >
                            Cancel
                        </Button>
                        <Button
                            type="submit"
                            className="flex-1"
                            isLoading={assignMutation.isPending}
                            disabled={!selectedAgentId || assignMutation.isPending}
                        >
                            Assign
                        </Button>
                    </div>
                </form>
            </Card>
        </div>
    );
};

export default AssignAgentModal;
