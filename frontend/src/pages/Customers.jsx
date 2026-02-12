import { DataTable } from '../components/DataTable';

const Customers = () => {
    const columns = [
        { key: 'id', label: 'ID' },
        { key: 'name', label: 'Name' },
        { key: 'email', label: 'Email' },
        { key: 'phone', label: 'Phone' },
        { key: 'company', label: 'Company' },
    ];

    const data = [
        { id: 'C001', name: 'Alice Wonderland', email: 'alice@wonderland.com', phone: '+1234567890', company: 'Wonderland Inc' },
        { id: 'C002', name: 'Mad Hatter', email: 'hatter@tea.com', phone: '+0987654321', company: 'Tea Party Co' },
    ];

    return (
        <div className="p-6">
            <h1 className="text-3xl font-bold text-gray-800 mb-6">Customers</h1>
            <DataTable columns={columns} data={data} title="Customer Database" onEdit={(row) => console.log('Edit', row)} />
        </div>
    );
};

export default Customers;
