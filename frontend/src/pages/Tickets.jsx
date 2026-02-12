import { useQuery, useMutation, useQueryClient } from '@tanstack/react-query';
import { useState } from 'react';
import { DataTable } from '../components/DataTable';
import api from '../api/axios';
import { format } from 'date-fns';

const Tickets = () => {
    const [page, setPage] = useState(1);
    const limit = 10;
    const queryClient = useQueryClient();

    const updateStatusMutation = useMutation({
        mutationFn: async ({ id, status }) => {
            const response = await api.put(`/tickets/${id}/status`, { status });
            return response.data;
        },
        onSuccess: () => {
            queryClient.invalidateQueries(['tickets']);
            alert('Ticket status updated successfully');
        },
        onError: (err) => {
            alert(`Failed to update status: ${err.response?.data?.error || err.message}`);
        },
    });

    const handleUpdateStatus = (id, status) => {
        if (window.confirm(`Are you sure you want to change status to ${status}?`)) {
            updateStatusMutation.mutate({ id, status });
        }
    };

    const columns = [
        { key: 'id', label: 'ID' },
        { key: 'title', label: 'Title' },
        {
            key: 'status',
            label: 'Status',
            render: (value) => (
                <span className={`px-2 py-1 rounded-full text-xs font-semibold ${value === 'open' ? 'bg-green-100 text-green-800' :
                    value === 'in_progress' ? 'bg-blue-100 text-blue-800' :
                        value === 'resolved' ? 'bg-gray-100 text-gray-800' :
                            value === 'requested' ? 'bg-purple-100 text-purple-800' :
                                'bg-yellow-100 text-yellow-800'
                    }`}>
                    {value ? value.replace('_', ' ').toUpperCase() : '-'}
                </span>
            )
        },
        {
            key: 'priority',
            label: 'Priority',
            render: (value) => (
                <span className={`px-2 py-1 rounded-full text-xs font-semibold ${value === 'high' ? 'bg-red-100 text-red-800' :
                    value === 'medium' ? 'bg-orange-100 text-orange-800' :
                        'bg-green-100 text-green-800'
                    }`}>
                    {value ? value.toUpperCase() : '-'}
                </span>
            )
        },
        {
            key: 'created_at',
            label: 'Created At',
            render: (value) => {
                if (!value) return '-';
                try {
                    return format(new Date(value), 'dd MMM yyyy HH:mm');
                } catch (e) {
                    return 'Invalid Date';
                }
            }
        },
    ];

    const { data, isLoading, error } = useQuery({
        queryKey: ['tickets', page, limit],
        queryFn: async () => {
            const response = await api.get(`/tickets?page=${page}&limit=${limit}`);
            return response.data;
        },
        keepPreviousData: true,
    });

    const tickets = data?.data || [];
    const meta = data?.meta || { total_pages: 1 };

    const renderActions = (ticket) => {
        if (ticket.status === 'requested') {
            return (
                <div className="flex gap-2">
                    <button
                        onClick={() => handleUpdateStatus(ticket.id, 'in_progress')}
                        className="text-xs bg-green-100 text-green-800 px-2 py-1 rounded hover:bg-green-200"
                    >
                        Approve
                    </button>
                    <button
                        onClick={() => handleUpdateStatus(ticket.id, 'closed')}
                        className="text-xs bg-red-100 text-red-800 px-2 py-1 rounded hover:bg-red-200"
                    >
                        Close
                    </button>
                </div>
            );
        }
        if (ticket.status === 'in_progress') {
            return (
                <div className="flex gap-2">
                    <button
                        onClick={() => handleUpdateStatus(ticket.id, 'resolved')}
                        className="text-xs bg-blue-100 text-blue-800 px-2 py-1 rounded hover:bg-blue-200"
                    >
                        Resolve
                    </button>
                    <button
                        onClick={() => handleUpdateStatus(ticket.id, 'closed')}
                        className="text-xs bg-red-100 text-red-800 px-2 py-1 rounded hover:bg-red-200"
                    >
                        Close
                    </button>
                </div>
            );
        }
        return null;
    };


    if (isLoading) return <div className="p-6">Loading...</div>;
    if (error) return <div className="p-6 text-red-500">Error loading tickets: {error.message}</div>;

    return (
        <div className="p-6">
            <h1 className="text-3xl font-bold text-gray-800 mb-6">Tickets</h1>
            <DataTable
                columns={columns}
                data={tickets}
                title="Support Tickets"
                actions={renderActions}
                currentPage={page}
                totalPages={meta.total_pages}
                onPageChange={setPage}
            />
        </div>
    );
};

export default Tickets;
