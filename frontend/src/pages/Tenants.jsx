import { DataTable } from '../components/DataTable';

const Tenants = () => {
    const columns = [
        { key: 'id', label: 'ID' },
        { key: 'name', label: 'Tenant Name' },
        { key: 'domain', label: 'Domain' },
        { key: 'plan', label: 'Plan' },
    ];

    const data = [
        { id: 'T001', name: 'Acme Corp', domain: 'acme.sociomile.com', plan: 'Enterprise' },
        { id: 'T002', name: 'Beta Ltd', domain: 'beta.sociomile.com', plan: 'Pro' },
        { id: 'T003', name: 'Gamma Inc', domain: 'gamma.sociomile.com', plan: 'Basic' },
    ];

    return (
        <div className="p-6">
            <h1 className="text-3xl font-bold text-gray-800 mb-6">Tenants</h1>
            <DataTable columns={columns} data={data} title="Registered Tenants" onEdit={(row) => console.log('Edit', row)} />
        </div>
    );
};

export default Tenants;
