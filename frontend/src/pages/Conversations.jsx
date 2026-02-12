import { DataTable } from '../components/DataTable';

const Conversations = () => {
    const columns = [
        { key: 'id', label: 'ID' },
        { key: 'customer', label: 'Customer' },
        { key: 'channel', label: 'Channel' },
        { key: 'lastMessage', label: 'Last Message' },
        { key: 'status', label: 'Status' },
    ];

    const data = [
        { id: '101', customer: 'Bob Brown', channel: 'WhatsApp', lastMessage: 'Hello, need help!', status: 'Open' },
        { id: '102', customer: 'Charlie Davis', channel: 'Email', lastMessage: 'Refund request', status: 'Pending' },
        { id: '103', customer: 'Eve White', channel: 'Live Chat', lastMessage: 'Thank you!', status: 'Closed' },
    ];

    return (
        <div className="p-6">
            <h1 className="text-3xl font-bold text-gray-800 mb-6">Conversations</h1>
            <DataTable columns={columns} data={data} title="Recent Conversations" actions={(row) => <button className="text-purple-600">View</button>} />
        </div>
    );
};

export default Conversations;
