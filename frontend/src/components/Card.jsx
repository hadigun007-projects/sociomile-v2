import { clsx } from "clsx";
import { twMerge } from "tailwind-merge";

export function Card({ className, children, ...props }) {
    return (
        <div
            className={twMerge(
                "bg-white rounded-lg shadow-md p-6 sm:p-8",
                className
            )}
            {...props}
        >
            {children}
        </div>
    );
}
