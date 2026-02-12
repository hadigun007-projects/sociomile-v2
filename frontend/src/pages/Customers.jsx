import { useQuery } from '@tanstack/react-query';
import { useState } from 'react';
import { DataTable } from '../components/DataTable';
import api from '../api/axios';
import { format } from 'date-fns';

const Customers = () => {
    const [page, setPage] = useState(1);
    const limit = 10;

    const columns = [
        { key: 'ExternalID', label: 'External ID' },
        { key: 'Name', label: 'Name' },
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
        queryKey: ['customers', page, limit],
        queryFn: async () => {
            const response = await api.get(`/customers?page=${page}&limit=${limit}`);
            return response.data;
        },
        keepPreviousData: true,
    });

    const customers = data?.data || [];
    const meta = data?.meta || { total_pages: 1 };

    if (isLoading) return <div className="p-6">Loading...</div>;
    if (error) return <div className="p-6 text-red-500">Error loading customers: {error.message}</div>;

    return (
        <div className="p-6">
            <h1 className="text-3xl font-bold text-gray-800 mb-6">Customers</h1>
            <DataTable
                columns={columns}
                data={customers}
                title="Customer Database"
                onEdit={(row) => console.log('Edit', row)}
                currentPage={page}
                totalPages={meta.total_pages}
                onPageChange={setPage}
            />
        </div>
    );
};

export default Customers;
