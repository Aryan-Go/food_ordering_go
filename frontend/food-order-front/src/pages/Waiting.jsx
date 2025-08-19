import React, { useState, useEffect } from "react";
import axios from "../axios/AxiosHandler.jsx";
import { toast, Bounce } from "react-toastify";
import { useNavigate, useSearchParams } from "react-router-dom";
import Navbar from "../components/Navbar_customer.jsx";

const Waiting = () => {
  const navigate = useNavigate()
  const [data, set_data] = useState(null);
  const [searchParams, setSearchParams] = useSearchParams();
  const [items_req, set_item_req] = useState([]);
  const id = searchParams.get("order_id");
  useEffect(() => {
    set_data(null)
    const data_sent = {
      id: Number(id),
    };
    const get_data = async () => {
      const resp = await axios.post("/customer/render_waiting", data_sent);
      set_data(resp.data);
    };
    get_data();
  }, []);

  if (data == null || data == undefined) {
    // navigate(`/payment?order_id=${id}`)
    return (
      <>
        {/* <Navbar /> */}
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
            <h1 className="text-2xl text-red-500">There is no data for this order id</h1>
          </div>
        </form>
      </>
    );
    navigate(`/payment?order_id=${id}`);
  } else {
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
      navigate(`/payment?order_id=${id}`);
    }
    else {
      return (
        <>
          {/* <Navbar /> */}
          <form action="/payment" method="post" className="text-center">
            <div className="h-[100vh] w-[100vw] bg-[url(waiting_back2.png)] backdrop-brightness-150 bg-no-repeat bg-size-[length:100%_100%] text-center flex flex-col items-center gap-2rem text-center">
              <h1 className="text-[5rem] pb-[2rem] font-bold">Ordered Items</h1>
              <table className="text-center w-[70vw] m-[4rem] px-[10rem] backdrop-blur-xs">
                <thead>
                  <tr>
                    <th scope="col" className="text-3xl mx-[2rem]">
                      Food Items
                    </th>
                    <th scope="col" className="text-3xl mx-[2rem]">
                      Special Instructions
                    </th>
                    <th scope="col" className="text-3xl mx-[2rem]">
                      Quantity Added
                    </th>
                    <th scope="col" className="text-3xl mx-[2rem]">
                      Prepared/left
                    </th>
                  </tr>
                </thead>
                <tbody>
                  {data.map((value, key) => (
                    <tr key={key}>
                      <td className="text-2xl font-bold mx-[2rem]">
                        {value.food_name}
                      </td>
                      <td className="text-2xl font-semibold mx-[2rem]">
                        {value.instructions == "" ? "None" : value.instructions}
                      </td>
                      <td className="text-2xl font-bold mx-[2rem]">
                        {value.quant}
                      </td>
                      <td className="text-2xl font-bold mx-[2rem]">
                        {value.status}
                      </td>
                    </tr>
                  ))}
                </tbody>
              </table>
            </div>
          </form>
        </>
      );
    }
    }
    
};
export default Waiting;
