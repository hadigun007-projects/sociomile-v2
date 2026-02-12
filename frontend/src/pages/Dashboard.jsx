import { FaUsers, FaChartLine, FaWallet, FaEllipsisV } from 'react-icons/fa';

const Dashboard = () => {
    return (
        <div className="pt-6">
            <header className="mb-8">
                <h1 className="text-4xl font-normal text-gray-900 mb-2">Dashboard</h1>
                <p className="text-gray-600 text-lg">Detailed overview of your platform</p>
            </header>

            {/* Dashboard Widgets */}
            <div className="grid grid-cols-1 md:grid-cols-3 gap-6 mb-8">
                {/* Card 1 */}
                <div className="group relative overflow-hidden bg-white rounded-[28px] p-8 shadow-sm hover:shadow-md transition-all duration-300 border border-purple-50">
                    <div className="absolute top-0 right-0 p-8 opacity-10 group-hover:scale-110 transition-transform duration-500">
                        <FaUsers className="text-9xl text-purple-600" />
                    </div>
                    <div className="relative z-10">
                        <div className="w-14 h-14 rounded-2xl bg-purple-100 text-purple-700 flex items-center justify-center text-2xl mb-6 group-hover:bg-purple-600 group-hover:text-white transition-colors duration-300">
                            <FaUsers />
                        </div>
                        <p className="text-gray-500 font-medium text-sm tracking-wide uppercase">Total Users</p>
                        <h3 className="text-4xl font-normal text-gray-900 mt-2">1,234</h3>
                        <div className="flex items-center mt-4 text-sm font-medium text-green-600 bg-green-50 w-fit px-3 py-1 rounded-full">
                            <span>+12.5%</span>
                            <span className="text-gray-500 ml-2 font-normal">vs last month</span>
                        </div>
                    </div>
                </div>

                {/* Card 2 */}
                <div className="group relative overflow-hidden bg-white rounded-[28px] p-8 shadow-sm hover:shadow-md transition-all duration-300 border border-purple-50">
                    <div className="absolute top-0 right-0 p-8 opacity-10 group-hover:scale-110 transition-transform duration-500">
                        <FaChartLine className="text-9xl text-purple-600" />
                    </div>
                    <div className="relative z-10">
                        <div className="w-14 h-14 rounded-2xl bg-purple-100 text-purple-700 flex items-center justify-center text-2xl mb-6 group-hover:bg-purple-600 group-hover:text-white transition-colors duration-300">
                            <FaChartLine />
                        </div>
                        <p className="text-gray-500 font-medium text-sm tracking-wide uppercase">Active Sessions</p>
                        <h3 className="text-4xl font-normal text-gray-900 mt-2">567</h3>
                        <div className="flex items-center mt-4 text-sm font-medium text-green-600 bg-green-50 w-fit px-3 py-1 rounded-full">
                            <span>+5.2%</span>
                            <span className="text-gray-500 ml-2 font-normal">vs last week</span>
                        </div>
                    </div>
                </div>

                {/* Card 3 */}
                <div className="group relative overflow-hidden bg-white rounded-[28px] p-8 shadow-sm hover:shadow-md transition-all duration-300 bg-purple-600 text-white">
                    <div className="absolute top-0 right-0 p-8 opacity-20 group-hover:scale-110 transition-transform duration-500">
                        <FaWallet className="text-9xl text-white" />
                    </div>
                    <div className="relative z-10">
                        <div className="w-14 h-14 rounded-2xl bg-white/20 text-white flex items-center justify-center text-2xl mb-6 backdrop-blur-sm">
                            <FaWallet />
                        </div>
                        <p className="text-purple-100 font-medium text-sm tracking-wide uppercase">Total Revenue</p>
                        <h3 className="text-4xl font-normal mt-2">$12,345</h3>
                        <div className="flex items-center mt-4 text-sm font-medium text-white bg-white/20 w-fit px-3 py-1 rounded-full backdrop-blur-md">
                            <span>+18.2%</span>
                            <span className="text-purple-100 ml-2 font-normal">vs last month</span>
                        </div>
                    </div>
                </div>
            </div>

            <div className="grid grid-cols-1 lg:grid-cols-3 gap-8">
                {/* Recent Activity */}
                <div className="lg:col-span-2 bg-white rounded-[28px] p-8 border border-white/50 shadow-sm">
                    <div className="flex justify-between items-center mb-8">
                        <div>
                            <h3 className="text-2xl font-normal text-gray-900">Recent Activity</h3>
                            <p className="text-gray-500 text-sm mt-1">Latest actions across the platform</p>
                        </div>
                        <button className="p-2 rounded-full hover:bg-gray-100 text-gray-500">
                            <FaEllipsisV />
                        </button>
                    </div>

                    <div className="space-y-2">
                        {[1, 2, 3, 4].map((i) => (
                            <div key={i} className="group flex items-center p-4 rounded-2xl hover:bg-purple-50 transition-colors cursor-pointer">
                                <div className={`w-12 h-12 rounded-xl flex items-center justify-center text-lg font-bold mr-5 ${i % 2 === 0 ? 'bg-purple-100 text-purple-700' : 'bg-pink-100 text-pink-700'
                                    }`}>
                                    U{i}
                                </div>
                                <div className="flex-1">
                                    <div className="flex justify-between">
                                        <h4 className="font-semibold text-gray-800 text-base">User Activity {i}</h4>
                                        <span className="text-xs text-gray-400 font-medium">2h ago</span>
                                    </div>
                                    <p className="text-sm text-gray-500 mt-1 line-clamp-1">Updated project settings and deployed new changes to production.</p>
                                </div>
                                <div className="ml-4 opacity-0 group-hover:opacity-100 transition-opacity">
                                    <span className="text-xs font-bold text-purple-600 bg-purple-100 px-3 py-1 rounded-full">View</span>
                                </div>
                            </div>
                        ))}
                    </div>
                    <div className="mt-6 pt-4 border-t border-gray-100 text-center">
                        <button className="text-purple-600 font-semibold hover:text-purple-800 text-sm py-2 px-4 rounded-full hover:bg-purple-50 transition-colors">
                            View All Activity
                        </button>
                    </div>
                </div>

                {/* System Status / Right Panel */}
                <div className="bg-white rounded-[28px] p-8 border border-white/50 shadow-sm flex flex-col h-full">
                    <h3 className="text-2xl font-normal text-gray-900 mb-6">System Health</h3>

                    <div className="space-y-8 flex-1">
                        <div className="relative p-6 bg-purple-50 rounded-3xl overflow-hidden">
                            <div className="flex justify-between items-end mb-2 relative z-10">
                                <span className="text-gray-600 font-medium">Server Load</span>
                                <span className="text-3xl font-normal text-purple-700">45%</span>
                            </div>
                            <div className="w-full bg-purple-200/50 rounded-full h-3 overflow-hidden relative z-10">
                                <div className="bg-purple-600 h-3 rounded-full" style={{ width: '45%' }}></div>
                            </div>
                            <div className="absolute right-0 bottom-0 opacity-10 transform translate-x-1/4 translate-y-1/4">
                                <div className="w-32 h-32 bg-purple-600 rounded-full"></div>
                            </div>
                        </div>

                        <div className="relative p-6 bg-pink-50 rounded-3xl overflow-hidden">
                            <div className="flex justify-between items-end mb-2 relative z-10">
                                <span className="text-gray-600 font-medium">Memory Usage</span>
                                <span className="text-3xl font-normal text-pink-700">60%</span>
                            </div>
                            <div className="w-full bg-pink-200/50 rounded-full h-3 overflow-hidden relative z-10">
                                <div className="bg-pink-600 h-3 rounded-full" style={{ width: '60%' }}></div>
                            </div>
                            <div className="absolute right-0 bottom-0 opacity-10 transform translate-x-1/4 translate-y-1/4">
                                <div className="w-32 h-32 bg-pink-600 rounded-full"></div>
                            </div>
                        </div>

                        <div className="relative p-6 bg-indigo-50 rounded-3xl overflow-hidden">
                            <div className="flex justify-between items-end mb-2 relative z-10">
                                <span className="text-gray-600 font-medium">DB Connections</span>
                                <span className="text-3xl font-normal text-indigo-700">28%</span>
                            </div>
                            <div className="w-full bg-indigo-200/50 rounded-full h-3 overflow-hidden relative z-10">
                                <div className="bg-indigo-600 h-3 rounded-full" style={{ width: '28%' }}></div>
                            </div>
                            <div className="absolute right-0 bottom-0 opacity-10 transform translate-x-1/4 translate-y-1/4">
                                <div className="w-32 h-32 bg-indigo-600 rounded-full"></div>
                            </div>
                        </div>
                    </div>
                </div>
            </div>
        </div>
    );
};

export default Dashboard;
