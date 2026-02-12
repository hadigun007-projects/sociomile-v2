import { useQuery, useMutation, useQueryClient } from '@tanstack/react-query';
import { DataTable } from '../components/DataTable';
import api from '../api/axios';
import { useState } from 'react';
import { PlusIcon } from '@heroicons/react/24/outline';
import AddTenantModal from '../components/AddTenantModal';
import EditTenantModal from '../components/EditTenantModal';
import { Card } from '../components/Card';

const Tenants = () => {
    const [selectedTenant, setSelectedTenant] = useState(null);
    const [isCreateModalOpen, setIsCreateModalOpen] = useState(false);
    const [isEditModalOpen, setIsEditModalOpen] = useState(false);
    const [page, setPage] = useState(1);
    const limit = 10;
    const queryClient = useQueryClient();

    const currentUser = JSON.parse(localStorage.getItem('user') || '{}');
    const isSuperAdmin = currentUser.role === 'super_admin';

    const { data, isLoading } = useQuery({
        queryKey: ['tenants', page, limit],
        queryFn: async () => {
            const response = await api.get(`/tenants?page=${page}&limit=${limit}`);
            return response.data;
        },
        enabled: isSuperAdmin,
        keepPreviousData: true,
    });

    const tenants = data?.data || [];
    const meta = data?.meta || { total_pages: 1 };

    const deleteMutation = useMutation({
        mutationFn: async (id) => {
            await api.delete(`/tenants/${id}`);
        },
        onSuccess: () => {
            queryClient.invalidateQueries(['tenants']);
        },
    });

    const handleEdit = (tenant) => {
        setSelectedTenant(tenant);
        setIsEditModalOpen(true);
    };

    const handleDelete = async (tenant) => {
        if (window.confirm('Are you sure you want to delete this tenant?')) {
            try {
                await deleteMutation.mutateAsync(tenant.id);
            } catch (error) {
                console.error('Error deleting tenant:', error);
                alert('Failed to delete tenant');
            }
        }
    };

    const columns = [
        { key: 'name', label: 'Name' },
        {
            key: 'plan',
            label: 'Plan',
            render: (value) => (
                <span className={`px-2 py-1 rounded-full text-xs font-semibold
                    ${value === 'enterprise' ? 'bg-purple-100 text-purple-700' :
                        value === 'pro' ? 'bg-blue-100 text-blue-700' : 'bg-gray-100 text-gray-700'}`}>
                    {value ? value.charAt(0).toUpperCase() + value.slice(1) : 'Basic'}
                </span>
            )
        },
        { key: 'created_at', label: 'Created At' },
    ];

    if (!isSuperAdmin) {
        return <div className="p-6 text-red-600">Access Denied</div>;
    }

    if (isLoading) return <div>Loading...</div>;

    return (
        <div className="space-y-6">
            <div className="flex justify-between items-center">
                <div>
                    <h1 className="text-3xl font-bold text-gray-800">Tenants</h1>
                    <p className="text-gray-600 mt-1">Manage organizations and subscriptions</p>
                </div>
                <button
                    onClick={() => setIsCreateModalOpen(true)}
                    className="bg-purple-600 text-white px-4 py-2 rounded-lg hover:bg-purple-700 transition-colors"
                >
                    Add Tenant
                </button>
            </div>

            <Card>
                <DataTable
                    columns={columns}
                    data={tenants}
                    onEdit={handleEdit}
                    onDelete={handleDelete}
                    emptyMessage="No tenants found"
                    currentPage={page}
                    totalPages={meta.total_pages}
                    onPageChange={setPage}
                />
            </Card>
            {selectedTenant && (
                <EditTenantModal
                    isOpen={isEditModalOpen}
                    onClose={() => setIsEditModalOpen(false)}
                    tenant={selectedTenant}
                />
            )}
            <AddTenantModal
                isOpen={isCreateModalOpen}
                onClose={() => setIsCreateModalOpen(false)}
            />
        </div>
    );
};

export default Tenants;
