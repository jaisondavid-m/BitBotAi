import React,{useState} from 'react'

function StudentHome() {
  const [question, setQuestion] = useState("");
  const [answer, setAnswer] = useState("");

    const askQuestion = async () => {
    const res = await fetch("http://localhost:8000/ask", {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify({ question }),
    });

    const data = await res.json();
    setAnswer(data.answer);
  };

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
