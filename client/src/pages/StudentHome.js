import React,{useState} from 'react'
import { api } from '../api/axios';
import NavBar from '../components/NavBar';
import { LuArrowUp } from "react-icons/lu";
import { CiStop1 } from "react-icons/ci";

function StudentHome() {
  const [question, setQuestion] = useState("");
  const [showquestion, setShowquestion] = useState("");
  const [answer, setAnswer] = useState("");
  const [loading,setLoading]=useState(false)

    const askQuestion = async(e)=>{
      e.preventDefault();
      try {
        setQuestion("")
        setLoading(true)
        setShowquestion(question)
        const res = await api.post("/ask",{question})
        setAnswer(res.data.answer)
      } catch (err) {
        console.error(err);
        setAnswer(err.response?.data?.error || "Something went wrong");
      }
      finally{
        setLoading(false)
      }
    }

  return (
    <div className="min-h-screen bg-black text-white flex flex-col">
      <NavBar />
      <h3 className="text-4xl font-bold mt-16  mb-2 text-center  ">{showquestion}</h3>
      <div className="flex-1 px-4 py-6 overflow-y-auto">
        {loading ? (<div className="animate-pulse space-y-3 mt-5 lg:w-[50%] mx-auto">
                        <div className="h-4 bg-gray-700 rounded w-3/4"></div>
                        <div className="h-4 bg-gray-700 rounded w-2/4"></div>
                        <div className="h-4 bg-gray-700 rounded w-full"></div>
                        <div className="h-4 bg-gray-700 rounded w-5/6"></div>
                    </div>
      ) : (
        <p className="lg:w-[50%] mx-auto text-mono text-2xl text-gray-300 pb-20"><pre className="whitespace-pre-wrap break-words">{answer}</pre></p>
      )}
      </div>
      <div className="fixed bottom-0 left-0 w-full bg-black pb-3 px-4">
        <form onSubmit={askQuestion} className="relative max-w-2xl mx-auto flex items-center">
          <input type="text" value={question} required placeholder="Ask your question..." onChange={(e) => setQuestion(e.target.value)}
            className="w-full h-12 rounded-full px-4 bg-[#1b1b1e] placeholder-gray-500"/>
          <button type='submit' className="absolute right-0.5 scale-105 bg-white  p-3 rounded-full transition">
            {loading?<CiStop1 size={20} className='text-black'/>:<LuArrowUp className="text-black" size={20}/>}
          </button>
        </form>
      </div>
    </div>

  )
}

export default StudentHome
