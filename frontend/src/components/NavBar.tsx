import React from "react";
import NavItem, { NavItemProps } from "./NavItem";
import "./style/NavBar.css"

import '../index.css'

interface NavBarProps {
  navItems: NavItemProps[];
  onCartClick: () => void;
}

const NavBar: React.FC<NavBarProps> = ({ navItems }) => {
  return (
    <nav className="nav-bar">
      <ul>
        {navItems.map((item, idx) => (
          <NavItem key={idx} {...item} />
        ))}
      </ul>
    </nav>
  );
};

export default NavBar;
