import { useState } from 'react';
import { Link, useNavigate } from 'react-router-dom';
import { useMutation } from '@tanstack/react-query';
import { FaFacebook, FaGoogle } from 'react-icons/fa';
import api from '../api/axios';
import { Button } from '../components/Button';
import Input from '../components/Input';
import { Card } from '../components/Card';

const Login = () => {
    const [email, setEmail] = useState('');
    const [password, setPassword] = useState('');
    const [errorMsg, setErrorMsg] = useState('');
    const navigate = useNavigate();

    const loginMutation = useMutation({
        mutationFn: async (credentials) => {
            const response = await api.post('/auth/login', credentials);
            return response.data;
        },
        onSuccess: (data) => {
            localStorage.setItem('token', data.access_token);
            localStorage.setItem('refreshToken', data.refresh_token);
            localStorage.setItem('user', JSON.stringify(data.user));

            navigate('/dashboard');
        },
        onError: (err) => {
            setErrorMsg(err.response?.data?.error || 'Login failed. Please check your credentials.');
        },
    });

    const handleSubmit = (e) => {
        e.preventDefault();
        setErrorMsg('');
        loginMutation.mutate({ email, password });
    };

    return (
        <div className="min-h-screen bg-surfaceVariant/20 flex items-center justify-center p-4">
            <Card className="w-full max-w-md bg-white">
                <div className="text-center mb-8">
                    <h2 className="text-3xl font-bold text-onSurface">Welcome Back</h2>
                    <p className="text-secondary mt-2">Please sign in to continue</p>
                </div>

                {errorMsg && (
                    <div className="bg-error/10 text-error p-3 rounded-lg mb-4 text-sm font-medium">
                        {errorMsg}
                    </div>
                )}

                <form onSubmit={handleSubmit} className="space-y-4">
                    <Input
                        label="Email"
                        type="email"
                        value={email}
                        onChange={(e) => setEmail(e.target.value)}
                        required
                        placeholder="Enter your email"
                    />
                    <Input
                        label="Password"
                        type="password"
                        value={password}
                        onChange={(e) => setPassword(e.target.value)}
                        required
                        placeholder="Enter your password"
                    />

                    <Button type="submit" className="w-full" isLoading={loginMutation.isPending}>
                        Log In
                    </Button>
                </form>
            </Card>
        </div>
    );
};

export default Login;
