import React, { forwardRef } from 'react';

const Input = forwardRef(({
    label,
    type = 'text',
    error,
    helperText,
    className = '',
    ...props
}, ref) => {
    const id = props.id || props.name || Math.random().toString(36).substr(2, 9);

    return (
        <div className={`relative mb-5 ${className}`}>
            <div className="relative">
                <input
                    ref={ref}
                    type={type}
                    id={id}
                    placeholder=" "
                    className={`
                        peer block w-full rounded-t-lg border-b-2 border-gray-300 bg-gray-50 px-2.5 pb-2.5 pt-5 text-sm text-gray-900 
                        focus:border-purple-600 focus:outline-none focus:ring-0 
                        disabled:opacity-50 disabled:cursor-not-allowed
                        transition-colors duration-200
                        ${error ? 'border-red-500 focus:border-red-500' : ''}
                    `}
                    {...props}
                />
                <label
                    htmlFor={id}
                    className={`
                        absolute left-2.5 top-4 z-10 origin-[0] -translate-y-4 scale-75 transform text-sm text-gray-500 duration-300 
                        peer-placeholder-shown:translate-y-0 peer-placeholder-shown:scale-100 
                        peer-focus:-translate-y-4 peer-focus:scale-75 peer-focus:text-purple-600
                        ${error ? 'text-red-500 peer-focus:text-red-500' : ''}
                    `}
                >
                    {label}
                </label>
            </div>
            {(error || helperText) && (
                <p className={`mt-1 text-xs ${error ? 'text-red-500' : 'text-gray-500'}`}>
                    {error || helperText}
                </p>
            )}
        </div>
    );
});

Input.displayName = 'Input';

export default Input;
