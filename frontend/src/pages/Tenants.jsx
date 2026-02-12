import { useQuery, useMutation, useQueryClient } from '@tanstack/react-query';
import { DataTable } from '../components/DataTable';
import api from '../api/axios';
import { useState } from 'react';
import { PlusIcon } from '@heroicons/react/24/outline';
import AddTenantModal from '../components/AddTenantModal';
import EditTenantModal from '../components/EditTenantModal';

const Tenants = () => {
    const [selectedTenant, setSelectedTenant] = useState(null);
    const [isEditModalOpen, setIsEditModalOpen] = useState(false);
    const [isAddModalOpen, setIsAddModalOpen] = useState(false);
    const queryClient = useQueryClient();

    const { data: tenants = [], isLoading, error } = useQuery({
        queryKey: ['tenants'],
        queryFn: async () => {
            const response = await api.get('/tenants');
            return response.data.data;
        }
    });

    const deleteMutation = useMutation({
        mutationFn: async (id) => {
            await api.delete(`/tenants/${id}`);
        },
        onSuccess: () => {
            queryClient.invalidateQueries(['tenants']);
        },
        onError: (err) => {
            alert(`Failed to delete tenant: ${err.message}`);
        }
    });

    const columns = [
        { key: 'id', label: 'ID' },
        { key: 'name', label: 'Tenant Name' },
        { key: 'plan', label: 'Plan' },
    ];

    const handleEdit = (tenant) => {
        setSelectedTenant(tenant);
        setIsEditModalOpen(true);
    };

    const handleDelete = (tenant) => {
        if (window.confirm(`Are you sure you want to delete tenant ${tenant.name}?`)) {
            deleteMutation.mutate(tenant.id);
        }
    };

    if (isLoading) return <div className="p-6">Loading...</div>;
    if (error) return <div className="p-6 text-red-500">Error loading tenants: {error.message}</div>;

    return (
        <div className="p-6">
            <div className="flex justify-between items-center mb-6">
                <h1 className="text-3xl font-bold text-gray-800">Tenants</h1>
                <button
                    onClick={() => setIsAddModalOpen(true)}
                    className="inline-flex items-center rounded-md bg-purple-600 px-3 py-2 text-sm font-semibold text-white shadow-sm hover:bg-purple-500 focus-visible:outline focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-purple-600"
                >
                    <PlusIcon className="-ml-0.5 mr-1.5 h-5 w-5" aria-hidden="true" />
                    Add Tenant
                </button>
            </div>
            <DataTable
                columns={columns}
                data={tenants}
                title="Registered Tenants"
                onEdit={handleEdit}
                onDelete={handleDelete}
            />
            {selectedTenant && (
                <EditTenantModal
                    isOpen={isEditModalOpen}
                    onClose={() => setIsEditModalOpen(false)}
                    tenant={selectedTenant}
                />
            )}
            <AddTenantModal
                isOpen={isAddModalOpen}
                onClose={() => setIsAddModalOpen(false)}
            />
        </div>
    );
};

export default Tenants;
