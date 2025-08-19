import React, { useEffect, useState } from "react";
import Navbar from "../components/Navbar_customer.jsx";
import axios from "../axios/AxiosHandler.jsx";
import { toast, Bounce } from "react-toastify";
import { useNavigate } from "react-router-dom";

const Customer = () => {
   const delay = (ms) => new Promise((resolve) => setTimeout(resolve, ms));
   const navigate = useNavigate();
   const [data, set_data] = useState([]);
   useEffect(() => {
     const authent = async () => {
       const resp = await axios.get("/customer");
       set_data(resp.data);
     };
     authent();
   }, []);
  if (data.status_code == 202) {
    
    return (
      <>
        <Navbar />
        <div className="min-h-[100vh] h-max bg-[url(customer_home2.webp)] bg-black/70 bg-blend-overlay bg-no-repeat bg-size-[length:100vw_130vh] px-[5rem]">
          <h1 className="font-bold mb-4 text-white dark:text-gray-600  text-center">
            <span className="text-6xl bg-clip-text text-yellow-300">
              Welcome To SERVXPRESS
            </span>
          </h1>
          <h1 className="font-bold text-center mb-4 text-white dark:text-white">
            <span className="text-3xl text-transparent bg-clip-text text-yellow-300 text-center">
              Better Food Best Services.
            </span>
          </h1>
          <p className="font-bold text-xl text-orange-500 dark:text-white text-shadow-black">
            Here at ServeXpress our motive is to provide world classNameic food
            and best services to each customer.
          </p>
          <h1 className="font-bold text-center mb-4 text-7xl text-white dark:text-white md:text-9xl lg:text-9xl">
            <span className="text-transparent bg-clip-text text-yellow-300 text-5xl text-center">
              Who are we
            </span>
          </h1>
          <div className="grid grid-cols-1 md:gap-6 md:grid-cols-2">
            <p className="font-bold mb-3 text-shadow-black text-white  text-2xl">
              This is an initiative started by ServXpress to improve service
              delivery and customer engagement. By leveraging innovative
              technology and streamlined operations, the company aims to provide
              faster, more reliable, and personalized experiences for its users.
              This step marks a commitment to continuous improvement and
              excellence in service.{" "}
            </p>
            <blockquote className="mb-3">
              <p className="font-bold text-2xl text-shadow-black italic text-orange-900 dark:text-white">
                " ServeXpress is just awesome. It offers a seamless, intuitive
                experience with ordering food from menu to michelin star chef.
                It's the perfect choice for your next online food ordering and
                delivery. "
              </p>
            </blockquote>
          </div>
          <h1 className="font-bold text-center mb-4 text-9xl  text-white dark:text-white md:text-9xl lg:text-9xl">
            <span className="text-transparent bg-clip-text text-yellow-300 text-5xl text-center">
              Different Perspectives
            </span>
          </h1>
          <div className="grid grid-cols-1 gap-6 sm:grid-cols-2">
            <p className="font-bold mb-3 text-white dark:text-white text-2xl text-shadow-black">
              👨‍🍳 As a Customer: At ServeXpress, customers enjoy a seamless and
              personalized food ordering experience. With a wide variety of
              cuisines, easy-to-use interface, real-time tracking, and reliable
              delivery, users can order their favorite meals with just a few
              taps. Whether it's a quick lunch or a family dinner, ServeXpress
              ensures quality food and timely service—every time.
            </p>
            <p className="font-bold mb-3 text-white dark:text-white text-2xl text-shadow-black">
              🏪 As a Dealer (Chef/Restaurant): ServeXpress empowers local chefs
              and restaurants by providing them a digital platform to reach more
              customers, manage orders efficiently, and grow their business.
              From order notifications to performance analytics and customer
              feedback, dealers get all the tools they need to deliver
              excellence while focusing on what they do best—cooking great food.
            </p>
          </div>
        </div>
      </>
    );
  }
  else {
      const show_message = async() => {
        toast.error(data.message, {
          position: "top-center",
          autoClose: 5000,
          hideProgressBar: false,
          closeOnClick: false,
          pauseOnHover: true,
          draggable: true,
          progress: undefined,
          theme: "dark",
          transition: Bounce,
        });
        await delay(3000);
        // navigate(-1)
      }
      show_message()
   }
};

export default Customer;
