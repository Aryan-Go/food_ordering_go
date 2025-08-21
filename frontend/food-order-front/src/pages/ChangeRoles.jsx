import React,{useEffect,useState} from 'react'
import axios from "../axios/AxiosHandler.jsx";
import Navbar from "../components/Navbar_admin.jsx";
import { toast, Bounce } from "react-toastify";

const ChangeRoles = () => {
    const [data,set_data] = useState([])
    const [customer_chef,set_customer_chef] = useState([])
    useEffect(() => {
        const get_data = async () => {
            let response = await axios.get("/admin/signup")
            set_data(response.data)
        }
        get_data()
        const get_data2 = async () => {
          const response = await axios.get("/admin/admin_details");
          set_customer_chef(response.data);
        };
        get_data2();
    } , [])
    console.log("This is the data");
    console.log(data);
    console.log(typeof data);
    const handleChef = async (id) => {
        const data_sent = {
          id: id,
        };
        const response = await axios.post(
          "/admin/admin_chef_conversion",
          data_sent
        );
        let response2 = await axios.get("/admin/signup");
        set_data(response2.data);
        if (data.status_code == 401) {
          toast.error(response.data.message, {
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
          if (response.data.message == "Malformed Token") {
            setTimeout(() => {
              navigate("/login");
            }, 2000);
          }
        } else {
          toast.success(response.data.message, {
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
        }
    }
    const handleAdmin = async (id) => {
      const data_sent = {
        id: id,
      };
      const response = await axios.post(
        "/admin/admin_conversion",
        data_sent
      );
        let response2 = await axios.get("/admin/signup");
        set_data(response2.data);
      if (data.status_code == 401) {
        toast.error(response.data.message, {
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
        if (response.data.message == "Malformed Token") {
          setTimeout(() => {
            navigate("/login");
          }, 2000);
        }
      } else {
        toast.success(response.data.message, {
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
      }
    };
    const user_table = data.map((value, key) => {
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
              value={ customer_chef.chef_customer == null ? "None" :
                customer_chef.chef_customer.includes(value.id) ? "chef" : "None"
              }
              name="food_id"
              className="food_id text-center"
            />
          </td>
          <td className="text-xl font-bold mx-[2rem] mt-[20rem]">
                  <br />
            <div className="flex flex-row">
            <button type="button" className="bg-white w-[12rem] m-[1rem] p-[1rem] rounded-2xl" onClick={() => handleChef(value.id)}>
              Change role to chef
            </button>
            <button type="button" className="bg-white w-[12rem] m-[1rem] p-[1rem] rounded-2xl" onClick={() => handleAdmin(value.id)}>
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
      <div className="h-full w-full bg-[url(user_background.jpg)] bg-no-repeat bg-size-[length:100%_100%] flex flex-col justify-center items-center gap-2rem text-center">
        <h1 className="text-[6rem] font-bold">User details</h1>
        <table className="text-center w-[70vw] m-[4rem] px-[10rem] table-auto">
          <thead>
            <tr className="m-[1rem]">
              <th scope="col" className="text-4xl mx-[2rem]">
                Name
              </th>
              <th scope="col" className="text-4xl mx-[2rem]">
                Email
              </th>
              <th scope="col" className="text-4xl mx-[2rem]">
                Role
              </th>
              <th scope="col" className="text-4xl mx-[2rem]">
                Requested Role
              </th>
              <th scope="col" className="text-4xl mx-[2rem]">
                Options
              </th>
            </tr>
          </thead>
          <tbody>{user_table}</tbody>
        </table>
      </div>
    </div>
  );
}

export default ChangeRoles