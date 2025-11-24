import React,{useState} from 'react'
import { api } from '../api/axios';

function StudentHome() {
  const [question, setQuestion] = useState("");
  const [answer, setAnswer] = useState("");

    const askQuestion = async(e)=>{
      e.preventDefault();
      try {
        const res = await api.post("/ask",{question})
        setAnswer(res.data.answer)
      } catch (err) {
        console.error(err);
        setAnswer(err.response?.data?.error || "Something went wrong");
      }
    }

  return (
    <div>
      <h2>Student Ask Question</h2>
      <input type="text" value={question} placeholder="Ask your question..." onChange={(e) => setQuestion(e.target.value)}/>
      <button onClick={askQuestion}>Ask</button>
      <h3>Answer</h3>
      <p style={{width:"500px"}}>{answer}</p>
    </div>
  )
}

export default StudentHome
