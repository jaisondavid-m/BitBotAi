import { useState } from "react"
import {useNavigate,Link} from "react-router-dom"
import { api } from "../api/axios"

function Login() {

    const [email,setEmail]=useState("")
    const [password,setPassword]=useState("")
    const [loading, setLoading] = useState(false);
    const navigate = useNavigate()

    const handleLogin = async (e)=>{
        e.preventDefault();
        setLoading(true)
        try {
            const res = await api.post("/login",{email,password})
            localStorage.setItem("token",res.data.token)
            navigate("/")
        } catch (err) {
            if (err.response && err.response.data?.error) {
                alert(err.response.data.error);
            } else {
                alert("Something went wrong");
            }
        }finally{
            setLoading(false)
        }
    }

  return (
   <div className='flex justify-center items-center bg-[#D78FEE] h-screen text-center'>
        <div className=' bg-[#9B5DE0] rounded-2xl h-max w-max p-10'>
          <form onSubmit={handleLogin} className='flex flex-col gap-y-5 text-center w-max'>
            <input className='p-3 rounded-xl' type='email' value={email} placeholder='Enter Your Email' onChange={(e)=>setEmail(e.target.value)}/>
            <input className='p-3 rounded-xl' type='password' value={password} placeholder='Enter Your Password' onChange={(e)=>setPassword(e.target.value)}/>
            <button className='bg-purple-950 text-white font-bold w-max px-4 py-2 rounded-xl mx-auto' type='submit'>{loading?"Logging In .." :"LogIn"}</button>
          </form>
          <div className='mt-5 text-white flex flex-col gap-y-2'>
            <p>Did not Have an account ?</p>
            <Link to="/register" className='text-purple-950 font-bold underline'>Create an Account</Link>
          </div>
        </div>
    </div>
  )
}

export default Login
