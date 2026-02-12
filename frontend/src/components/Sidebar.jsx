import { Link, useLocation } from 'react-router-dom';
import {
    FaUsers, FaComments, FaBuilding, FaTicketAlt,
    FaUserFriends, FaSignOutAlt, FaChartPie
} from 'react-icons/fa';

const SidebarItem = ({ icon: Icon, label, to }) => {
    const location = useLocation();
    const active = location.pathname === to || (to !== '/dashboard' && location.pathname.startsWith(to));

    return (
        <Link
            to={to}
            className={`flex items-center gap-4 px-6 py-3.5 mx-3 rounded-[20px] transition-all duration-300 group ${active
                ? 'bg-purple-100 text-purple-800 font-semibold shadow-sm'
                : 'text-gray-600 hover:bg-purple-50 hover:text-purple-700'
                }`}
        >
            <Icon className={`text-xl ${active ? 'text-purple-600' : 'text-gray-400 group-hover:text-purple-500'}`} />
            <span className="text-sm tracking-wide">{label}</span>
        </Link>
    );
};

export const Sidebar = ({ isSidebarOpen, user, handleLogout }) => {
    const menuItems = [
        { icon: FaChartPie, label: 'Dashboard', to: '/dashboard', roles: ['admin', 'agent', 'owner'] },
        { icon: FaUsers, label: 'Users', to: '/dashboard/users', roles: ['admin'] }, // Owner only manages tenants, Admin manages users
        { icon: FaComments, label: 'Conversations', to: '/dashboard/conversations', roles: ['agent', 'admin'] },
        { icon: FaBuilding, label: 'Tenants', to: '/dashboard/tenants', roles: ['owner'] },
        { icon: FaTicketAlt, label: 'Tickets', to: '/dashboard/tickets', roles: ['admin', 'agent'] },
        { icon: FaUserFriends, label: 'Customers', to: '/dashboard/customers', roles: ['admin', 'agent'] },
    ];

    const filteredMenuItems = menuItems.filter(item => item.roles.includes(user.role));

    return (
        <aside
            className={`fixed inset-y-0 left-0 z-40 w-72 bg-white/90 backdrop-blur-xl border-r border-white/50 shadow-lg transition-transform duration-300 ease-in-out transform ${isSidebarOpen ? 'translate-x-0' : '-translate-x-full'
                } lg:translate-x-0 lg:static lg:bg-transparent lg:shadow-none lg:border-none lg:w-72 flex-shrink-0`}
        >
            <div className="h-full flex flex-col p-4">
                <div className="flex items-center px-6 py-5 mb-6">
                    <span className="text-2xl font-bold text-purple-700 tracking-tight">Sociomile</span>
                </div>

                <div className="flex-1 overflow-y-auto space-y-1 [&::-webkit-scrollbar]:hidden [-ms-overflow-style:'none'] [scrollbar-width:'none']">
                    <div className="px-6 pb-2">
                        <p className="text-xs font-bold text-gray-400 uppercase tracking-wider">Main Menu</p>
                    </div>
                    {filteredMenuItems.map((item) => (
                        <SidebarItem
                            key={item.label}
                            icon={item.icon}
                            label={item.label}
                            to={item.to}
                        />
                    ))}
                </div>

                <div className="p-4 mt-auto">
                    <div className="bg-purple-600 rounded-[24px] p-5 text-white shadow-xl shadow-purple-200 relative overflow-hidden">
                        <div className="relative z-10">
                            <div className="flex items-center gap-3 mb-3">
                                <div className="w-10 h-10 rounded-full bg-white/20 backdrop-blur-md border border-white/30 p-0.5">
                                    <img
                                        src={`https://ui-avatars.com/api/?name=${user.email || 'U'}&background=random&color=fff`}
                                        alt="Profile"
                                        className="w-full h-full rounded-full object-cover"
                                    />
                                </div>
                                <div className="overflow-hidden">
                                    <p className="text-sm font-bold truncate">{user.email?.split('@')[0] || 'User'}</p>
                                    <p className="text-xs text-purple-200 truncate">{user.email}</p>
                                </div>
                            </div>
                            <button
                                onClick={handleLogout}
                                className="w-full bg-white/10 hover:bg-white/20 backdrop-blur-md text-white text-xs font-bold py-2.5 rounded-xl transition-all flex items-center justify-center gap-2"
                            >
                                <FaSignOutAlt />
                                Sign Out
                            </button>
                        </div>
                        {/* Decorative circles */}
                        <div className="absolute -top-6 -right-6 w-24 h-24 bg-white/10 rounded-full blur-xl"></div>
                        <div className="absolute -bottom-6 -left-6 w-24 h-24 bg-purple-500/30 rounded-full blur-xl"></div>
                    </div>
                </div>
            </div>
        </aside>
    );
};
