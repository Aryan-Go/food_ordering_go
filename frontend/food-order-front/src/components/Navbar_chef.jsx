import React from "react";
import { NavLink } from "react-router-dom";

const Navbar = () => {
  return (
    <nav className="bg-[#e3f2fd] " data-bs-theme="light">
      <div className="container-fluid flex flex-row justify-between">
        <NavLink className="navbar-brand" to="/chef">
          <img className="w-[7rem]" src="logo-removebg.png" alt="logo" />
        </NavLink>
        <div className="flex justify-around items-center w-[20%]" role="search">
          <NavLink
            className="text-[1.2rem] navbar-brand mb-0 h1 text-gray-600 no-underline font-bold"
            to="/order"
          >
            Orders
          </NavLink>
          <NavLink
            className="text-[1.2rem] navbar-brand mb-0 h1 text-gray-600 no-underline font-bold"
            to="/chef"
          >
            Home
          </NavLink>
          <NavLink href="/logout">
            <button type="button" className="text-[1.2rem] btn btn-primary font-bold bg-blue-400 p-[0.5rem] rounded-xl">
              Logout
            </button>
          </NavLink>
        </div>
      </div>
    </nav>
  );
};

export default Navbar;
