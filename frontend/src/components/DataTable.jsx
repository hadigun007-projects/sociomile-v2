import { useState, useMemo } from 'react';
import { Card } from './Card';
import { FaEdit, FaTrash, FaSearch } from 'react-icons/fa';

export const DataTable = ({ columns, data, title, actions, onDelete, onEdit }) => {
    const [searchTerm, setSearchTerm] = useState('');

    const filteredData = useMemo(() => {
        if (!searchTerm) return data;

        return data.filter(row =>
            Object.values(row).some(value =>
                String(value).toLowerCase().includes(searchTerm.toLowerCase())
            )
        );
    }, [data, searchTerm]);

    return (
        <Card className="p-0 overflow-hidden border-none shadow-sm">
            <div className="p-6 border-b border-gray-100 bg-white flex flex-col sm:flex-row justify-between items-center gap-4">
                {title && <h2 className="text-xl font-semibold text-gray-800">{title}</h2>}

                <div className="relative w-full sm:w-64">
                    <div className="absolute inset-y-0 left-0 pl-3 flex items-center pointer-events-none">
                        <FaSearch className="text-gray-400" />
                    </div>
                    <input
                        type="text"
                        placeholder="Search..."
                        className="block w-full pl-10 pr-3 py-2 border border-gray-200 rounded-full leading-5 bg-gray-50 placeholder-gray-400 focus:outline-none focus:bg-white focus:ring-1 focus:ring-purple-500 focus:border-purple-500 sm:text-sm transition-all"
                        value={searchTerm}
                        onChange={(e) => setSearchTerm(e.target.value)}
                    />
                </div>
            </div>

            <div className="overflow-x-auto">
                <table className="w-full text-left border-collapse">
                    <thead>
                        <tr className="bg-purple-50/50 text-purple-900 border-b border-purple-100">
                            {columns.map((col) => (
                                <th key={col.key} className="p-4 font-semibold text-sm tracking-wide">
                                    {col.label}
                                </th>
                            ))}
                            {(actions || onEdit || onDelete) && <th className="p-4 font-semibold text-sm tracking-wide text-right">Actions</th>}
                        </tr>
                    </thead>
                    <tbody className="divide-y divide-gray-50">
                        {filteredData.map((row, index) => (
                            <tr key={index} className="hover:bg-purple-50/30 transition-colors group">
                                {columns.map((col) => (
                                    <td key={col.key} className="p-4 text-sm text-gray-600 font-medium">
                                        {col.render ? col.render(row[col.key], row) : row[col.key]}
                                    </td>
                                ))}
                                {(actions || onEdit || onDelete) && (
                                    <td className="p-4 text-right">
                                        <div className="flex justify-end gap-2 opacity-0 group-hover:opacity-100 transition-opacity">
                                            {onEdit && (
                                                <button onClick={() => onEdit(row)} className="p-2 text-blue-600 hover:bg-blue-50 rounded-lg transition-colors">
                                                    <FaEdit />
                                                </button>
                                            )}
                                            {onDelete && (
                                                <button onClick={() => onDelete(row)} className="p-2 text-red-600 hover:bg-red-50 rounded-lg transition-colors">
                                                    <FaTrash />
                                                </button>
                                            )}
                                            {actions && actions(row)}
                                        </div>
                                    </td>
                                )}
                            </tr>
                        ))}
                        {filteredData.length === 0 && (
                            <tr>
                                <td colSpan={columns.length + (actions ? 1 : 0)} className="p-8 text-center text-gray-400 italic">
                                    {data.length === 0 ? "No data available" : "No matching records found"}
                                </td>
                            </tr>
                        )}
                    </tbody>
                </table>
            </div>
        </Card>
    );
};
