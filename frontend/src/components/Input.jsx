import { forwardRef } from "react";
import { clsx } from "clsx";
import { twMerge } from "tailwind-merge";

export const Input = forwardRef(({ label, className, id, ...props }, ref) => {
    return (
        <div className="mb-4">
            {label && (
                <label htmlFor={id || props.name} className="block text-sm font-medium text-gray-700 mb-1">
                    {label}
                </label>
            )}
            <input
                ref={ref}
                id={id || props.name}
                className={twMerge(
                    "appearance-none block w-full px-3 py-2 border border-outline rounded-md shadow-sm placeholder-gray-400 focus:outline-none focus:ring-primary focus:border-primary sm:text-sm",
                    className
                )}
                {...props}
            />
        </div>
    );
});

Input.displayName = "Input";
