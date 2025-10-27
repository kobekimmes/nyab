
import React from "react";
import '../index.css'
import "./style/NavItem.css"

export interface NavItemProps {
    href: string;
    icon?: string;
    display?: string;
    text?: string | React.ReactNode;
    external?: boolean;
    onClick?: () => void;
}


export const NavItem: React.FC<NavItemProps> = ({href, icon, display, text, external, onClick}) => {


    const content = (
        <>
            {(icon ? 
                <i className={icon}></i>
                : 
                <span>{display}</span>
            )}
            {(text ?
            <div className="hover-text">
                {text}
            </div>
            :
            <></>
            )}
        </>
    );

    return (
        <li className="nav-item">
            {onClick ? (
                <button
                onClick={onClick}
                className="nav-item-button"
                >
                    {content}
                </button>
            ) : (
                <a
                href={href}
                target={external ? "_blank" : undefined}
                rel={external ? "noopener noreferrer" : undefined}
                >
                    {content}
                </a>
            )}
        </li>
  );
}

export default NavItem;