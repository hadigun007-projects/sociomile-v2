import { useQuery } from '@tanstack/react-query';
import { DataTable } from '../components/DataTable';
import api from '../api/axios';

const Tenants = () => {
    const { data: tenants = [], isLoading, error } = useQuery({
        queryKey: ['tenants'],
        queryFn: async () => {
            const response = await api.get('/tenants');
            return response.data.data;
        }
    });

    const columns = [
        { key: 'id', label: 'ID' },
        { key: 'name', label: 'Tenant Name' },
        { key: 'plan', label: 'Plan' },
    ];

    if (isLoading) return <div className="p-6">Loading...</div>;
    if (error) return <div className="p-6 text-red-500">Error loading tenants: {error.message}</div>;

    return (
        <div className="p-6">
            <h1 className="text-3xl font-bold text-gray-800 mb-6">Tenants</h1>
            <DataTable columns={columns} data={tenants} title="Registered Tenants" onEdit={(row) => console.log('Edit', row)} />
        </div>
    );
};

export default Tenants;
