import { useQuery, useMutation, useQueryClient } from '@tanstack/react-query';
import { DataTable } from '../components/DataTable';
import api from '../api/axios';
import { useState } from 'react';
import EditUserModal from '../components/EditUserModal';
import AddUserModal from '../components/AddUserModal';
import { PlusIcon } from '@heroicons/react/24/outline';

const Users = () => {
    const [selectedUser, setSelectedUser] = useState(null);
    const [isEditModalOpen, setIsEditModalOpen] = useState(false);
    const [isAddModalOpen, setIsAddModalOpen] = useState(false);
    const queryClient = useQueryClient();

    const { data: users = [], isLoading, error } = useQuery({
        queryKey: ['users'],
        queryFn: async () => {
            const response = await api.get('/users');
            return response.data.data;
        }
    });

    const deleteMutation = useMutation({
        mutationFn: async (id) => {
            await api.delete(`/users/${id}`);
        },
        onSuccess: () => {
            queryClient.invalidateQueries(['users']);
        },
        onError: (err) => {
            alert(`Failed to delete user: ${err.message}`);
        }
    });

    const columns = [
        { key: 'id', label: 'ID' },
        { key: 'email', label: 'Email' },
        { key: 'role', label: 'Role' },
    ];

    const currentUser = JSON.parse(localStorage.getItem('user') || '{}');

    const handleEdit = (user) => {
        if (user.id === currentUser.id) return;
        setSelectedUser(user);
        setIsEditModalOpen(true);
    };

    const handleDelete = (user) => {
        if (user.id === currentUser.id) return;
        if (window.confirm(`Are you sure you want to delete user ${user.email}?`)) {
            deleteMutation.mutate(user.id);
        }
    };

    const isActionDisabled = (user) => user.id === currentUser.id;

    if (isLoading) return <div className="p-6">Loading...</div>;
    if (error) return <div className="p-6 text-red-500">Error loading users: {error.message}</div>;

    return (
        <div className="p-6">
            <div className="flex justify-between items-center mb-6">
                <h1 className="text-3xl font-bold text-gray-800">Users</h1>
                <button
                    onClick={() => setIsAddModalOpen(true)}
                    className="inline-flex items-center rounded-md bg-purple-600 px-3 py-2 text-sm font-semibold text-white shadow-sm hover:bg-purple-500 focus-visible:outline focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-purple-600"
                >
                    <PlusIcon className="-ml-0.5 mr-1.5 h-5 w-5" aria-hidden="true" />
                    Add User
                </button>
            </div>
            <DataTable
                columns={columns}
                data={users}
                title="System Users"
                onEdit={handleEdit}
                onDelete={handleDelete}
                isActionDisabled={isActionDisabled}
            />
            {selectedUser && (
                <EditUserModal
                    isOpen={isEditModalOpen}
                    onClose={() => setIsEditModalOpen(false)}
                    user={selectedUser}
                />
            )}
            <AddUserModal
                isOpen={isAddModalOpen}
                onClose={() => setIsAddModalOpen(false)}
            />
        </div>
    );
};

export default Users;
