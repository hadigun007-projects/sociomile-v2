import { Fragment, useState } from 'react';
import { Dialog, Transition } from '@headlessui/react';
import { XMarkIcon } from '@heroicons/react/24/outline';
import api from '../api/axios';
import { useMutation, useQueryClient } from '@tanstack/react-query';
import Input from './Input';

const AddTenantModal = ({ isOpen, onClose }) => {
    const [name, setName] = useState('');
    const [plan, setPlan] = useState('basic');
    const [adminEmail, setAdminEmail] = useState('');
    const [adminPassword, setAdminPassword] = useState('');
    const queryClient = useQueryClient();

    const createMutation = useMutation({
        mutationFn: async (data) => {
            const response = await api.post('/tenants', data);
            return response.data;
        },
        onSuccess: () => {
            queryClient.invalidateQueries(['tenants']);
            onClose();
            // Reset form
            setName('');
            setPlan('basic');
            setAdminEmail('');
            setAdminPassword('');
        },
    });

    const handleSubmit = async (e) => {
        e.preventDefault();
        const data = {
            name,
            plan,
            admin_email: adminEmail,
            admin_password: adminPassword
        };
        createMutation.mutate(data);
    };

    return (
        <Transition.Root show={isOpen} as={Fragment}>
            <Dialog as="div" className="relative z-50" onClose={onClose}>
                <Transition.Child
                    as={Fragment}
                    enter="ease-out duration-300"
                    enterFrom="opacity-0"
                    enterTo="opacity-100"
                    leave="ease-in duration-200"
                    leaveFrom="opacity-100"
                    leaveTo="opacity-0"
                >
                    <div className="fixed inset-0 bg-gray-500 bg-opacity-75 transition-opacity" />
                </Transition.Child>

                <div className="fixed inset-0 z-10 overflow-y-auto">
                    <div className="flex min-h-full items-end justify-center p-4 text-center sm:items-center sm:p-0">
                        <Transition.Child
                            as={Fragment}
                            enter="ease-out duration-300"
                            enterFrom="opacity-0 translate-y-4 sm:translate-y-0 sm:scale-95"
                            enterTo="opacity-100 translate-y-0 sm:scale-100"
                            leave="ease-in duration-200"
                            leaveFrom="opacity-100 translate-y-0 sm:scale-100"
                            leaveTo="opacity-0 translate-y-4 sm:translate-y-0 sm:scale-95"
                        >
                            <Dialog.Panel className="relative transform overflow-hidden rounded-lg bg-white px-4 pb-4 pt-5 text-left shadow-xl transition-all sm:my-8 sm:w-full sm:max-w-lg sm:p-6">
                                <div className="absolute right-0 top-0 hidden pr-4 pt-4 sm:block">
                                    <button
                                        type="button"
                                        className="rounded-md bg-white text-gray-400 hover:text-gray-500 focus:outline-none focus:ring-2 focus:ring-purple-500 focus:ring-offset-2"
                                        onClick={onClose}
                                    >
                                        <span className="sr-only">Close</span>
                                        <XMarkIcon className="h-6 w-6" aria-hidden="true" />
                                    </button>
                                </div>
                                <div className="sm:flex sm:items-start">
                                    <div className="mt-3 text-center sm:ml-4 sm:mt-0 sm:text-left w-full">
                                        <Dialog.Title as="h3" className="text-base font-semibold leading-6 text-gray-900 mb-6">
                                            Add New Tenant
                                        </Dialog.Title>
                                        <div className="mt-2">
                                            <form onSubmit={handleSubmit} className="space-y-4">
                                                <Input
                                                    label="Tenant Name"
                                                    type="text"
                                                    id="name"
                                                    value={name}
                                                    onChange={(e) => setName(e.target.value)}
                                                    required
                                                />
                                                <Input
                                                    label="Admin Email"
                                                    type="email"
                                                    id="admin_email"
                                                    value={adminEmail}
                                                    onChange={(e) => setAdminEmail(e.target.value)}
                                                    required
                                                />
                                                <Input
                                                    label="Admin Password"
                                                    type="password"
                                                    id="admin_password"
                                                    value={adminPassword}
                                                    onChange={(e) => setAdminPassword(e.target.value)}
                                                    required
                                                    minLength={6}
                                                />
                                                <div>
                                                    <label htmlFor="plan" className="block text-sm font-medium text-gray-700 mb-1">
                                                        Plan
                                                    </label>
                                                    <select
                                                        id="plan"
                                                        name="plan"
                                                        className="block w-full rounded-t-lg border-b-2 border-gray-300 bg-gray-50 px-2.5 pb-2.5 pt-2.5 text-sm text-gray-900 focus:border-purple-600 focus:outline-none focus:ring-0"
                                                        value={plan}
                                                        onChange={(e) => setPlan(e.target.value)}
                                                    >
                                                        <option value="basic">Basic</option>
                                                        <option value="premium">Premium</option>
                                                        <option value="enterprise">Enterprise</option>
                                                    </select>
                                                </div>
                                                <div className="mt-5 sm:mt-4 sm:flex sm:flex-row-reverse">
                                                    <button
                                                        type="submit"
                                                        className="inline-flex w-full justify-center rounded-md bg-purple-600 px-3 py-2 text-sm font-semibold text-white shadow-sm hover:bg-purple-500 sm:ml-3 sm:w-auto"
                                                    >
                                                        Add Tenant
                                                    </button>
                                                    <button
                                                        type="button"
                                                        className="mt-3 inline-flex w-full justify-center rounded-md bg-white px-3 py-2 text-sm font-semibold text-gray-900 shadow-sm ring-1 ring-inset ring-gray-300 hover:bg-gray-50 sm:mt-0 sm:w-auto"
                                                        onClick={onClose}
                                                    >
                                                        Cancel
                                                    </button>
                                                </div>
                                            </form>
                                        </div>
                                    </div>
                                </div>
                            </Dialog.Panel>
                        </Transition.Child>
                    </div>
                </div>
            </Dialog>
        </Transition.Root>
    );
};

export default AddTenantModal;
