import React, { useState } from 'react'
import axios from "../axios/AxiosHandler.jsx"
import {toast,Bounce} from "react-toastify"
import { useNavigate } from "react-router-dom"

const Signup = () => {
    const [name,SetName] = useState("")
    const [email,SetEmail] = useState("")
    const [password,SetPassword] = useState("")
    const [repassword, SetRePassword] = useState("")
    const role = "customer"
    const navigate = useNavigate()
    const clickHandler = async(name, email, password, repassword, role) => {
        const data = {
            name: name,
            email: email,
            password: password,
            repassword: repassword,
            role: role
        };
        const response = await axios.post("/signup", data)
        if (response.data.status_code == 400) {
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
            })
        }
        else {
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
            })
            navigate("/login")
        }
    }
  return (
      <div className="w-[100vw] h-[100vh] bg-size-[100%_100%] grayscale-30 bg-[url(https://imgs.search.brave.com/TouscCUt008bVNGiF-OauhIv-4eqphkTGUWPWSRV25g/rs:fit:860:0:0:0/g:ce/aHR0cHM6Ly9pbWcu/ZnJlZXBpay5jb20v/cHJlbWl1bS12ZWN0/b3Ivd29ybGQtZm9v/ZC1kYXktYmFja2dy/b3VuZF83MjYyMzct/Mjk4LmpwZz9zZW10/PWFpc19oeWJyaWQm/dz03NDA)] bg-no-repeat bg-center flex justify-center items-center">
          <div className="w-max h-[37rem] p-[2rem] rounded-2xl bg-[#372A36] rounded-2xl">
              <h1 className="text-6xl font-bold bg-gradient-to-r from-blue-600 via-green-500 to-indigo-400 inline-block text-transparent bg-clip-text">Welcome customer to Xpress</h1>
              <form className="flex flex-col justify-center items-center" onSubmit={(e) => e.preventDefault()}>
                  <label htmlFor="name" className="text-white font-bold text-[1.5rem]">Name</label>
                  <input name="name" type="text" id="name" className="text-white w-[20rem] font-medium outline-none border-2 rounded-2xl m-[1rem] pl-[0.5rem] p-[0.2rem]" value={name} onChange={(e) => SetName(e.target.value)} />
                  <label htmlFor="email" className="text-white font-bold text-[1.5rem]">Email</label>
                  <input name="email" type="email" id="email" className="text-white w-[20rem] font-medium outline-none border-2 rounded-2xl m-[1rem] pl-[0.5rem] p-[0.2rem]" value={email} onChange={(e) => SetEmail(e.target.value)} />
                  <label htmlFor="password" className="text-white font-bold text-[1.5rem]">Password</label>
                  <input name="password" type="password" id="passsword" className="text-white w-[20rem] font-medium outline-none border-2 rounded-2xl m-[1rem] pl-[0.5rem] p-[0.2rem]" value={password} onChange={(e) => SetPassword(e.target.value)} />
                  <label htmlFor="repassword" className="text-white font-bold text-[1.5rem]">Repassword</label>
                  <input name="repassword" type="password" id="repassword" className="text-white w-[20rem] font-medium outline-none border-2 rounded-2xl m-[1rem] pl-[0.5rem] p-[0.2rem]" value={repassword} onChange={(e) => SetRePassword(e.target.value)} />
                  <input readOnly name="role" hidden type="text" className="text-white w-[20rem] font-medium outline-none border-2 rounded-2xl m-[1rem] pl-[0.5rem] p-[0.2rem]" value={role} />
                  <div>
                    <button type="submit" className="bg-gray-500 px-[1rem] rounded-2xl text-white m-[1rem] text-bold text-[2rem]" onClick={() => clickHandler(name,email,password,repassword,role)}>Submit</button>
                    <button type="submit" className="bg-gray-500 px-[1rem] rounded-2xl text-white m-[1rem] text-bold text-[2rem]" onClick={() => navigate("/login")}>Login</button>
                      
                  </div>
              </form>
          </div>
    </div>
  )
}

export default Signup