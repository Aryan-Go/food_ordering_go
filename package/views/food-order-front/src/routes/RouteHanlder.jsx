import { Routes, Route } from "react-router-dom"
import Signup from "../pages/Signup.jsx"
import Login from "../pages/Login.jsx"
import Customer from "../pages/Customer.jsx"
import Chef from "../pages/Chef.jsx"
import Customer_chef from "../pages/Customer_chef.jsx"
import Menu from "../pages/Menu.jsx"
import Waiting from "../pages/Waiting.jsx"
import Order from "../pages/Order.jsx"
import Admin from "../pages/Admin.jsx"

const RouteHanlder = () => {
  return (
    <Routes>
      <Route path="/" element={<Signup />} />
      <Route path="/login" element={<Login />} />
      <Route path="/customer" element={<Customer />} />
      <Route path="/chef" element={<Chef />} />
      <Route path="/customer_chef" element={<Customer_chef />} />
      <Route path="/menu" element={<Menu />} />
      <Route path="/waiting_page" element={<Waiting />} />
      <Route path="/order" element={<Order />} />
      <Route path="/admin" element={<Admin />} />
    </Routes>
  );
    
}

export default RouteHanlder