import { DataTable } from '../components/DataTable';

const Users = () => {
    const columns = [
        { key: 'id', label: 'ID' },
        { key: 'name', label: 'Name' },
        { key: 'email', label: 'Email' },
        { key: 'role', label: 'Role' },
        { key: 'status', label: 'Status' },
    ];

    const data = [
        { id: '1', name: 'John Doe', email: 'john@example.com', role: 'Admin', status: 'Active' },
        { id: '2', name: 'Jane Smith', email: 'jane@example.com', role: 'User', status: 'Active' },
        { id: '3', name: 'Alice Johnson', email: 'alice@example.com', role: 'Manager', status: 'Inactive' },
    ];

    return (
        <div className="p-6">
            <h1 className="text-3xl font-bold text-gray-800 mb-6">Users Management</h1>
            <DataTable columns={columns} data={data} title="All Users" onDelete={(row) => console.log('Delete', row)} onEdit={(row) => console.log('Edit', row)} />
        </div>
    );
};

export default Users;
