import React,{useEffect,useState} from "react";
import Navbar from "../components/Navbar_chef.jsx";
import axios from "../axios/AxiosHandler.jsx";
import { toast, Bounce } from "react-toastify";
import {useNavigate} from "react-router-dom"

const Chef = () => {
  const delay = (ms) => new Promise((resolve) => setTimeout(resolve, ms));
  const navigate = useNavigate()
  const [data,set_data] = useState([])
  useEffect(() => {
    const authent = async () => {
      const resp = await axios.get("/chef")
      set_data(resp.data)
    }
    authent()
  }, [])
  if (data.status_code == 202) {
     return (
       <>
         <Navbar />
         <div class="min-h-[100vh] h-max bg-[url(chef_back.jpg)] bg-black/50 bg-blend-overlay bg-no-repeat bg-size-[length:100vw_120vh]">
           <h1 class="mb-4 font-extrabold text-gray-600 dark:text-gray-600  text-center">
             <span class="text-8xl text-transparent bg-clip-text bg-gradient-to-r to-purple-600 from-sky-400">
               Welcome To SERVXPRESS
             </span>
           </h1>

           <h1 class="text-center mb-4 font-extrabold text-gray-600 dark:text-gray-600">
             <span class="text-5xl text-transparent bg-clip-text bg-gradient-to-r to-purple-300 from-white text-center">
               Better Food Best Services.
             </span>
           </h1>
           <p class="text-2xl font-normal text-orange-500 dark:text-white text-shadow-black">
             Here at ServeXpress our motive is to provide world classic food and
             best services to each customer.
           </p>
           <h1 class="text-center mb-4 text-9xl font-extrabold text-gray-600 dark:text-gray-600 md:text-9xl lg:text-9xl">
             <span class="text-transparent bg-clip-text bg-gradient-to-r to-purple-300 from-white text-5xl text-center">
               Who are we
             </span>
           </h1>
           <div class="grid grid-cols-1 md:gap-6 md:grid-cols-2">
             <p class="mb-3 text-shadow-black text-white  text-2xl">
               This is an initiative started by ServXpress to improve service
               delivery and customer engagement. By leveraging innovative
               technology and streamlined operations, the company aims to
               provide faster, more reliable, and personalized experiences for
               its users. This step marks a commitment to continuous improvement
               and excellence in service.{" "}
             </p>
             <blockquote class="mb-3">
               <p class="text-2xl text-shadow-black italic font-semibold text-orange-900 dark:text-white">
                 " ServeXpress is just awesome. It offers a seamless, intuitive
                 experience with ordering food from menu to michelin star chef.
                 It's the perfect choice for your next online food ordering and
                 delivery. "
               </p>
             </blockquote>
           </div>
           <p class="mb-3 text-2xl text-shadow-black ext-orange-500 dark:text-white range-400">
             Deliver great service experiences fast - without the complexity of
             traditional delays.Get guarenteed results with delicious food and
             crazy servics and benifits.
           </p>
           <h1 class="text-center mb-4 text-9xl font-extrabold text-gray-600 dark:text-gray-600 md:text-9xl lg:text-9xl">
             <span class="text-transparent bg-clip-text bg-gradient-to-r to-purple-300 from-white text-5xl text-center">
               Different Perspectives
             </span>
           </h1>
           <div class="grid grid-cols-1 gap-6 sm:grid-cols-2">
             <p class="mb-3 text-gray-600 dark:text-white text-2xl text-shadow-black">
               👨‍🍳 As a Customer: At ServeXpress, customers enjoy a seamless and
               personalized food ordering experience. With a wide variety of
               cuisines, easy-to-use interface, real-time tracking, and reliable
               delivery, users can order their favorite meals with just a few
               taps. Whether it's a quick lunch or a family dinner, ServeXpress
               ensures quality food and timely service—every time.
             </p>
             <p class="mb-3 text-gray-600 dark:text-white text-2xl text-shadow-black">
               🏪 As a Dealer (Chef/Restaurant): ServeXpress empowers local
               chefs and restaurants by providing them a digital platform to
               reach more customers, manage orders efficiently, and grow their
               business. From order notifications to performance analytics and
               customer feedback, dealers get all the tools they need to deliver
               excellence while focusing on what they do best—cooking great
               food.
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

export default Chef;
