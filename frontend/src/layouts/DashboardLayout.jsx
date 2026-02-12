import { useState } from 'react';
import { Outlet, useNavigate } from 'react-router-dom';
import { Sidebar } from '../components/Sidebar';
import { FaBars, FaTimes, FaSearch, FaBell } from 'react-icons/fa';

const DashboardLayout = () => {
    const [isSidebarOpen, setIsSidebarOpen] = useState(true);
    const navigate = useNavigate();
    const userString = localStorage.getItem('user');
    const user = userString ? JSON.parse(userString) : {};

    const handleLogout = () => {
        localStorage.removeItem('token');
        localStorage.removeItem('refreshToken');
        localStorage.removeItem('user');
        navigate('/login');
    };

    return (
        <div className="h-screen overflow-hidden bg-[#F3E5F5] font-sans selection:bg-purple-200 flex">
            {/* Sidebar */}
            <Sidebar
                isSidebarOpen={isSidebarOpen}
                user={user}
                handleLogout={handleLogout}
            />

            {/* Main Content Area */}
            <div className="flex-1 flex flex-col min-w-0 overflow-hidden">
                {/* Mobile Header */}
                <header className="lg:hidden bg-white/80 backdrop-blur-md border-b border-gray-100 p-4 flex items-center justify-between sticky top-0 z-30">
                    <div className="flex items-center gap-3">
                        <button onClick={() => setIsSidebarOpen(!isSidebarOpen)} className="p-2 text-gray-600 rounded-lg hover:bg-gray-100">
                            {isSidebarOpen ? <FaTimes /> : <FaBars />}
                        </button>
                        <span className="font-bold text-lg text-gray-800">Sociomile</span>
                    </div>
                    <div className="w-8 h-8 rounded-full bg-purple-100 overflow-hidden">
                        <img
                            src={`https://ui-avatars.com/api/?name=${user.email || 'U'}&background=9C27B0&color=fff`}
                            alt="Profile"
                        />
                    </div>
                </header>

                <main className="flex-1 overflow-y-auto p-4 lg:p-8 scroll-smooth">
                    <Outlet />
                </main>
            </div>
        </div>
    );
};

export default DashboardLayout;
