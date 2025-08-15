import React from 'react'

const Navbar_admin = () => {
  return (
     <nav class="navbar sticky-top bg-body-tertiary flex flex-row justify-around items-center" style="background-color: #e3f2fd;" data-bs-theme="light">
        <div class="container-fluid flex flex-row justify-around items-center">
          <a class="navbar-brand"><img class="w-[5rem]" src="/image/logo-removebg.png" alt="logo" /></a>
          <div class="flex flex-row justify-around w-[30%] h-[3rem]" role="search">
            <a class="navbar-brand mb-0 h1 text-gray-600 no-underline" href="/admin">Home</a>
            <a href="/logout"><button type="button" class="btn btn-primary">Logout</button></a>
          </div>
        </div>
    </nav>
  )
}

export default Navbar_admin