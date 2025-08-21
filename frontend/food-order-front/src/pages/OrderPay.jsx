import React, { useEffect, useState } from "react";
import axios from "../axios/AxiosHandler.jsx";
import Navbar from "../components/Navbar_admin.jsx";
import { toast, Bounce } from "react-toastify";
const OrderPay = () => {
    const [incomplete_order,set_incomplete_order] = useState([])
    const [incomplete_payment, set_incomplete_payment] = useState([])
    useEffect(() => {
      const get_data = async () => {
        const response = await axios.get("/admin/admin_details");
        set_incomplete_order(response.data.incomplete_order)
        set_incomplete_payment(response.data.incomplete_payment);
      };
        get_data();
    }, []);
    // for (let i = 0; i < )
    const food_table = incomplete_order.map((value, key) => {
      return (
        <tr key={key} className="m-[2rem] p-[2rem] mb-[30rem]">
          <td className="text-2xl font-bold mx-[2rem] mb-[20rem]">
            {value.name}
          </td>
          <td className="text-xl font-semibold mx-[2rem] w-[25rem] mb-[30rem] p-[2rem]">
            {value.email}
          </td>
          <td className="text-xl font-bold mx-[2rem] p-[2rem] mb-[20rem]">
            {value.role}
          </td>
          <td className="text-xl font-bold mx-[2rem] text-center mb-[20rem]">
            <input
              readOnly
              type="text"
              value={
                customer_chef.chef_customer == null
                  ? "None"
                  : customer_chef.chef_customer.includes(value.id)
                  ? "chef"
                  : "None"
              }
              name="food_id"
              className="food_id text-center"
            />
          </td>
          <td className="text-xl font-bold mx-[2rem] mt-[20rem]">
            <br />
            <div className="flex flex-row">
              <button
                type="button"
                className="bg-white w-[12rem] m-[1rem] p-[1rem] rounded-2xl"
                onClick={() => handleChef(value.id)}
              >
                Change role to chef
              </button>
              <button
                type="button"
                className="bg-white w-[12rem] m-[1rem] p-[1rem] rounded-2xl"
                onClick={() => handleAdmin(value.id)}
              >
                Change role to Admin
              </button>
            </div>
          </td>
        </tr>
      );
    });
  return (
    <div>
      <Navbar />
      <div className="min-h-screen h-full w-full bg-[url(waiting_back2.png)] bg-no-repeat bg-size-[length:100%_100%] flex flex-col justify-center items-center gap-2rem text-center">
        <div className="backdrop-blur-xs">
          <h1 className="text-[6rem] font-bold">User details</h1>
          <table className="text-center w-[70vw] m-[4rem] px-[10rem] table-auto">
            <thead>
              <tr className="m-[1rem]">
                <th scope="col" className="text-4xl mx-[2rem]">
                  Order Id
                </th>
                <th scope="col" className="text-4xl mx-[2rem]">
                  Customer Name
                </th>
                <th scope="col" className="text-4xl mx-[2rem]">
                  Chef Name
                </th>
                <th scope="col" className="text-4xl mx-[2rem]">
                  Go to Order
                </th>
                <th scope="col" className="text-4xl mx-[2rem]">
                  Total Payment
                </th>
                <th scope="col" className="text-4xl mx-[2rem]">
                  Complete Payment
                </th>
              </tr>
            </thead>
            <tbody>{food_table}</tbody>
          </table>
        </div>
      </div>
    </div>
  );
};

export default OrderPay;
