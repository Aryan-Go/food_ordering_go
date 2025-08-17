import React, { useState, useEffect } from "react";
import axios from "../axios/AxiosHandler.jsx";
import { toast, Bounce } from "react-toastify";
import Navbar from "../components/Navbar_chef.jsx";

const Order = () => {
    let counter = 0;
  const [data, set_data] = useState(null);
  useEffect(() => {
    set_data(null);
    const get_data = async () => {
      const resp = await axios.get("/chef/render_order");
      set_data(resp.data);
    };
    get_data();
  }, []);
  console.log("This the result I am getting");

  if (data == null || data == undefined) {
    return (
      <>
        <Navbar />
        <form action="/payment" method="post" className="text-center">
          <div className="h-[100vh] w-[100vw] bg-[url(menu_back.webp)] bg-no-repeat bg-size-[length:100%_100%] text-center flex flex-col items-center gap-2rem text-center">
            <h1 className="text-[6rem] font-bold">Ordered Items</h1>
            <table className="text-center w-[60vw] m-[4rem] px-[10rem]">
              <thead>
                <tr>
                  <th scope="col" className="text-2xl mx-[2rem]">
                    Food Items
                  </th>
                  <th scope="col" className="text-2xl mx-[2rem]">
                    Special Instructions
                  </th>
                  <th scope="col" className="text-2xl mx-[2rem]">
                    Quantity Added
                  </th>
                  <th scope="col" className="text-2xl mx-[2rem]">
                    Prepared/left
                  </th>
                </tr>
              </thead>
            </table>
            <h1 className="text-red-500 text-5xl font-bold">
              Please login as chef
            </h1>
          </div>
        </form>
      </>
    );
  } else {
    console.log(data);
    if (data.status_code != undefined) {
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
      return (
        <>
          <Navbar />
          <form action="/payment" method="post" className="text-center">
            <div className="h-[100vh] w-[100vw] bg-[url(waiting_back.webp)] bg-no-repeat bg-size-[length:100%_100%] text-center flex flex-col items-center gap-2rem text-center">
              <h1 className="text-[6rem] font-bold">Ordered Items</h1>
              <table className="text-center w-[60vw] m-[4rem] px-[10rem]">
                <thead>
                  <tr>
                    <th scope="col" className="text-2xl mx-[2rem]">
                      Food Items
                    </th>
                    <th scope="col" className="text-2xl mx-[2rem]">
                      Special Instructions
                    </th>
                    <th scope="col" className="text-2xl mx-[2rem]">
                      Quantity Added
                    </th>
                    <th scope="col" className="text-2xl mx-[2rem]">
                      Prepared/left
                    </th>
                  </tr>
                </thead>
              </table>
              <h1 className="text-red-500 text-5xl font-bold">
                No items to be made
              </h1>
            </div>
          </form>
        </>
      );
    } else {
        const clickHanlder = async (order_id, food_id) => {
          const data_sent = {
            food_id: food_id,
            order_id: order_id,
          };
            await axios.post("/chef/complete_order", data_sent);
            counter++;
            const get_data = async () => {
              const resp = await axios.get("/chef/render_order");
              set_data(resp.data);
            };
            get_data();
        };
      return (
        <>
          <Navbar />
            <div className="h-[100vh] w-[100vw] bg-[url(waiting_back.webp)] bg-no-repeat bg-size-[length:100%_100%] text-center flex flex-col items-center gap-2rem text-center">
              <h1 className="text-[6rem] font-bold">Ordered Items</h1>
              <table className="text-center w-[60vw] m-[4rem] px-[10rem]">
                <thead>
                  <tr>
                    <th scope="col" className="text-2xl mx-[2rem]">
                      Food Items
                    </th>
                    <th scope="col" className="text-2xl mx-[2rem]">
                      Special Instructions
                    </th>
                    <th scope="col" className="text-2xl mx-[2rem]">
                      Quantity Added
                    </th>
                    <th scope="col" className="text-2xl mx-[2rem]">
                      Prepared/left
                    </th>
                  </tr>
                </thead>
                <tbody>
                  {data.map((value, key) => (
                    <tr key={key}>
                      <td className="text-xl font-bold mx-[2rem]">
                        {value.food_name}
                      </td>
                      <td className="text-xl font-semibold mx-[2rem]">
                        {value.instructions}
                      </td>
                      <td className="text-xl font-bold mx-[2rem]">
                        {value.quant}
                      </td>
                      <td className="text-xl font-bold mx-[2rem]">
                              <button
                                  type="button"
                                  className="text-white bg-blue-400 p-[1rem] rounded-2xl"
                          onClick={() => clickHanlder(value.order_id , value.food_id)}
                        >
                          Mark as Completed
                        </button>
                      </td>
                    </tr>
                  ))}
                </tbody>
              </table>
            </div>
        </>
      );
    }
  }
};
export default Order;
