import { useQuery, useMutation, useQueryClient } from '@tanstack/react-query';
import { DataTable } from '../components/DataTable';
import api from '../api/axios';
import { useState } from 'react';
import EditUserModal from '../components/EditUserModal';
import AddUserModal from '../components/AddUserModal';
import { PlusIcon } from '@heroicons/react/24/outline';
import { Card } from '../components/Card';

const Users = () => {
    const [selectedUser, setSelectedUser] = useState(null);
    const [isCreateModalOpen, setIsCreateModalOpen] = useState(false);
    const [isEditModalOpen, setIsEditModalOpen] = useState(false);
    const [page, setPage] = useState(1);
    const limit = 10;
    const queryClient = useQueryClient();

    const { data, isLoading } = useQuery({
        queryKey: ['users', page, limit],
        queryFn: async () => {
            const response = await api.get(`/users?page=${page}&limit=${limit}`);
            return response.data;
        },
        keepPreviousData: true,
    });

    const users = data?.data || [];
    const meta = data?.meta || { total_pages: 1 };

    const deleteMutation = useMutation({
        mutationFn: async (id) => {
            await api.delete(`/users/${id}`);
        },
        onSuccess: () => {
            queryClient.invalidateQueries(['users']);
        },
    });

    const handleDelete = async (user) => {
        if (window.confirm('Are you sure you want to delete this user?')) {
            try {
                await deleteMutation.mutateAsync(user.id);
            } catch (error) {
                console.error('Error deleting user:', error);
                alert('Failed to delete user');
            }
        }
    };

    const handleEdit = (user) => {
        setSelectedUser(user);
        setIsEditModalOpen(true);
    };

    const columns = [
        { key: 'email', label: 'Email' },
        {
            key: 'role',
            label: 'Role',
            render: (value) => (
                <span className={`px-2 py-1 rounded-full text-xs font-semibold
                    ${value === 'admin' ? 'bg-purple-100 text-purple-700' : 'bg-blue-100 text-blue-700'}`}>
                    {value.charAt(0).toUpperCase() + value.slice(1)}
                </span>
            )
        },
        { key: 'created_at', label: 'Created At' },
    ];

    const currentUser = JSON.parse(localStorage.getItem('user') || '{}');
    const isActionDisabled = (user) => user.id === currentUser.id;

    if (isLoading) return <div>Loading...</div>;

    return (
        <div className="space-y-6">
            <div className="flex justify-between items-center">
                <div>
                    <h1 className="text-3xl font-bold text-gray-800">Users</h1>
                    <p className="text-gray-600 mt-1">Manage system users and their roles</p>
                </div>
                <button
                    onClick={() => setIsCreateModalOpen(true)}
                    className="bg-purple-600 text-white px-4 py-2 rounded-lg hover:bg-purple-700 transition-colors"
                >
                    Add User
                </button>
            </div>

            <Card>
                <DataTable
                    columns={columns}
                    data={users}
                    onEdit={handleEdit}
                    onDelete={handleDelete}
                    emptyMessage="No users found"
                    currentPage={page}
                    totalPages={meta.total_pages}
                    onPageChange={setPage}
                    isActionDisabled={isActionDisabled}
                />
            </Card>

            {selectedUser && (
                <EditUserModal
                    isOpen={isEditModalOpen}
                    onClose={() => setIsEditModalOpen(false)}
                    user={selectedUser}
                />
            )}
            <AddUserModal
                isOpen={isCreateModalOpen}
                onClose={() => setIsCreateModalOpen(false)}
            />
        </div>
    );
};

export default Users;
