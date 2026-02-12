import { useState } from 'react';
import { useMutation, useQuery } from '@tanstack/react-query';
import api from '../api/axios';
import { Card } from '../components/Card';
import Input from '../components/Input';
import { Button } from '../components/Button';

const ChannelSimulator = () => {
    const [tenantId, setTenantId] = useState('');
    const [customerExternalId, setCustomerExternalId] = useState('');
    const [message, setMessage] = useState('');
    const [responseMessage, setResponseMessage] = useState('');

    // Fetch tenants for dropdown
    const { data: tenants = [], isLoading: tenantsLoading } = useQuery({
        queryKey: ['public-tenants'],
        queryFn: async () => {
            const response = await api.get('/tenants/public');
            return response.data.data;
        },
    });

    const webhookMutation = useMutation({
        mutationFn: async (data) => {
            const response = await api.post('/channel/webhook', data);
            return response.data;
        },
        onSuccess: (data) => {
            setResponseMessage(`Message sent successfully! Conversation created for customer ${data.data.customer_external_id}`);
            setMessage('');
        },
        onError: (err) => {
            setResponseMessage(`Error: ${err.response?.data?.error || err.message}`);
        },
    });

    const handleSubmit = (e) => {
        e.preventDefault();
        setResponseMessage('');
        webhookMutation.mutate({
            tenant_id: tenantId,
            customer_external_id: customerExternalId,
            message: message,
        });
    };

    return (
        <div className="min-h-screen bg-gradient-to-br from-purple-50 via-white to-purple-50 p-6">
            <div className="max-w-6xl mx-auto">
                <div className="mb-8 text-center">
                    <h1 className="text-4xl font-bold text-gray-800 mb-2">Channel Webhook Simulator</h1>
                    <p className="text-gray-600">
                        Simulate incoming messages from external channels (WhatsApp, Instagram, etc.)
                    </p>
                </div>

                <div className="grid grid-cols-1 lg:grid-cols-2 gap-6">
                    <Card className="bg-white">
                        <h2 className="text-xl font-semibold text-gray-800 mb-4">Send Test Message</h2>
                        <form onSubmit={handleSubmit} className="space-y-4">
                            <div>
                                <label className="block text-sm font-medium text-gray-700 mb-1">
                                    Tenant
                                </label>
                                <select
                                    value={tenantId}
                                    onChange={(e) => setTenantId(e.target.value)}
                                    required
                                    disabled={tenantsLoading}
                                    className="block w-full rounded-lg border border-gray-300 px-3 py-2 shadow-sm focus:border-purple-500 focus:outline-none focus:ring-1 focus:ring-purple-500"
                                >
                                    <option value="">Select a tenant...</option>
                                    {tenants.map((tenant) => (
                                        <option key={tenant.id} value={tenant.id}>
                                            {tenant.id} - {tenant.name}
                                        </option>
                                    ))}
                                </select>
                            </div>
                            <Input
                                label="Customer External ID"
                                type="text"
                                value={customerExternalId}
                                onChange={(e) => setCustomerExternalId(e.target.value)}
                                required
                                placeholder="e.g., +6281234567890"
                            />
                            <div>
                                <label className="block text-sm font-medium text-gray-700 mb-1">
                                    Message
                                </label>
                                <textarea
                                    value={message}
                                    onChange={(e) => setMessage(e.target.value)}
                                    required
                                    rows={4}
                                    className="block w-full rounded-lg border border-gray-300 px-3 py-2 shadow-sm focus:border-purple-500 focus:outline-none focus:ring-1 focus:ring-purple-500"
                                    placeholder="Enter customer message..."
                                />
                            </div>

                            <Button
                                type="submit"
                                className="w-full"
                                isLoading={webhookMutation.isPending}
                            >
                                Send Message
                            </Button>
                        </form>

                        {responseMessage && (
                            <div
                                className={`mt-4 p-3 rounded-lg text-sm font-medium ${responseMessage.startsWith('✅')
                                    ? 'bg-green-50 text-green-700'
                                    : 'bg-red-50 text-red-700'
                                    }`}
                            >
                                {responseMessage}
                            </div>
                        )}
                    </Card>

                    <Card className="bg-purple-50">
                        <h2 className="text-xl font-semibold text-gray-800 mb-4">How It Works</h2>
                        <div className="space-y-3 text-sm text-gray-700">
                            <div className="flex items-start">
                                <span className="inline-flex items-center justify-center h-6 w-6 rounded-full bg-purple-600 text-white font-bold mr-3 flex-shrink-0">
                                    1
                                </span>
                                <p>
                                    <strong>Customer sends message</strong> via external channel (WhatsApp, Instagram, etc.)
                                </p>
                            </div>
                            <div className="flex items-start">
                                <span className="inline-flex items-center justify-center h-6 w-6 rounded-full bg-purple-600 text-white font-bold mr-3 flex-shrink-0">
                                    2
                                </span>
                                <p>
                                    <strong>System finds or creates customer</strong> based on external ID
                                </p>
                            </div>
                            <div className="flex items-start">
                                <span className="inline-flex items-center justify-center h-6 w-6 rounded-full bg-purple-600 text-white font-bold mr-3 flex-shrink-0">
                                    3
                                </span>
                                <p>
                                    <strong>System finds or creates conversation</strong> (open status)
                                </p>
                            </div>
                            <div className="flex items-start">
                                <span className="inline-flex items-center justify-center h-6 w-6 rounded-full bg-purple-600 text-white font-bold mr-3 flex-shrink-0">
                                    4
                                </span>
                                <p>
                                    <strong>Message is saved</strong> to conversation with sender type 'customer'
                                </p>
                            </div>
                        </div>

                        <div className="mt-6 p-3 bg-white rounded-lg border border-purple-200">
                            <p className="text-xs text-gray-600">
                                <strong>Note:</strong> This simulator acts as an external webhook endpoint.
                                In production, this would be triggered by actual channel platforms.
                            </p>
                        </div>
                    </Card>
                </div>
            </div>
        </div>
    );
};

export default ChannelSimulator;
