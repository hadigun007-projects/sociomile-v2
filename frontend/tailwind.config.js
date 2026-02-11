/** @type {import('tailwindcss').Config} */
export default {
    content: [
        "./index.html",
        "./src/**/*.{js,ts,jsx,tsx}",
    ],
    theme: {
        extend: {
            colors: {
                primary: '#3b82f6', // blue-500
                secondary: '#6b7280', // gray-500
                error: '#ef4444', // red-500
                surfaceVariant: '#f3f4f6', // gray-100
                onSurface: '#1f2937', // gray-800
                outline: '#d1d5db', // gray-300
            },
        },
    },
    plugins: [],
}
