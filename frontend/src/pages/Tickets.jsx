import { DataTable } from '../components/DataTable';

const Tickets = () => {
    const columns = [
        { key: 'id', label: 'ID' },
        { key: 'subject', label: 'Subject' },
        { key: 'requester', label: 'Requester' },
        { key: 'priority', label: 'Priority' },
        { key: 'status', label: 'Status' },
    ];

    const data = [
        { id: 'TIC-1001', subject: 'Login Issue', requester: 'Bob Brown', priority: 'High', status: 'Open' },
        { id: 'TIC-1002', subject: 'Billing Question', requester: 'Charlie Davis', priority: 'Medium', status: 'In Progress' },
        { id: 'TIC-1003', subject: 'Feature Request', requester: 'Eve White', priority: 'Low', status: 'New' },
    ];

    return (
        <div className="p-6">
            <h1 className="text-3xl font-bold text-gray-800 mb-6">Tickets</h1>
            <DataTable columns={columns} data={data} title="Support Tickets" actions={(row) => <button className="text-purple-600">Resolve</button>} />
        </div>
    );
};

export default Tickets;
