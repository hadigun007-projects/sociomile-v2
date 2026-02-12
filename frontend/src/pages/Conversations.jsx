import { useQuery } from '@tanstack/react-query';
import api from '../api/axios';
import { DataTable } from '../components/DataTable';
import { Card } from '../components/Card';

const Conversations = () => {
    const { data: conversations = [], isLoading } = useQuery({
        queryKey: ['conversations'],
        queryFn: async () => {
            const response = await api.get('/conversations');
            return response.data.data;
        },
    });

    const columns = [
        { key: 'id', label: 'ID' },
        { key: 'customer_id', label: 'Customer ID' },
        { key: 'status', label: 'Status' },
        { key: 'assigned_agent_id', label: 'Assigned Agent' },
        { key: 'created_at', label: 'Created At' },
    ];

    if (isLoading) {
        return (
            <div className="flex items-center justify-center h-64">
                <div className="animate-spin rounded-full h-12 w-12 border-b-2 border-purple-600"></div>
            </div>
        );
    }

    return (
        <div className="space-y-6">
            <div className="flex items-center justify-between">
                <div>
                    <h1 className="text-3xl font-bold text-gray-800">Conversations</h1>
                    <p className="text-gray-600 mt-1">
                        Manage customer conversations and messages
                    </p>
                </div>
                <div className="flex gap-2">
                    <select className="px-4 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-purple-500">
                        <option value="">All Status</option>
                        <option value="open">Open</option>
                        <option value="assigned">Assigned</option>
                        <option value="closed">Closed</option>
                    </select>
                </div>
            </div>

            <Card>
                <DataTable
                    columns={columns}
                    data={conversations}
                    emptyMessage="No conversations found"
                />
            </Card>

            {/* Stats Cards */}
            <div className="grid grid-cols-1 md:grid-cols-4 gap-4">
                <Card className="bg-green-50">
                    <div className="text-sm text-green-600 font-medium">Open</div>
                    <div className="text-2xl font-bold text-green-700 mt-1">
                        {conversations.filter((c) => c.status === 'open').length}
                    </div>
                </Card>
                <Card className="bg-blue-50">
                    <div className="text-sm text-blue-600 font-medium">Assigned</div>
                    <div className="text-2xl font-bold text-blue-700 mt-1">
                        {conversations.filter((c) => c.status === 'assigned').length}
                    </div>
                </Card>
                <Card className="bg-gray-50">
                    <div className="text-sm text-gray-600 font-medium">Closed</div>
                    <div className="text-2xl font-bold text-gray-700 mt-1">
                        {conversations.filter((c) => c.status === 'closed').length}
                    </div>
                </Card>
                <Card className="bg-purple-50">
                    <div className="text-sm text-purple-600 font-medium">Total</div>
                    <div className="text-2xl font-bold text-purple-700 mt-1">
                        {conversations.length}
                    </div>
                </Card>
            </div>
        </div>
    );
};

export default Conversations;
