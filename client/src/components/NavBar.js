import React,{useState} from 'react'
import { SiVowpalwabbit } from "react-icons/si";
import { Link } from 'react-router-dom';
import Hamburger from 'hamburger-react'
import { motion, AnimatePresence } from "framer-motion";
import { SiGooglegemini } from "react-icons/si";    

function NavBar() {
  const [isOpen, setIsOpen] = useState(false);

  return (
    <>
      <div className='sticky top-0 flex justify-between w-screen py-2 px-5 md:p-10 lg:h-20 lg:px-36 lg:py-5 bg-white/5 backdrop-blur-md shadow-lg'>
        <div className="flex gap-x-2 lg:justify-between items-center">
          <SiVowpalwabbit className="w-7 h-10 md:w-12 md:h-12 lg:w-14 lg:h-14" />
          <h1 className="text-2xl lg:text-5xl mt-1 font-bold tracking-tight font-mono">BIT-AI</h1>
        </div>
        <div className='hidden lg:flex gap-8 font-bold text-2xl'>
          <p>PROFILE</p>
          <p>ABOUT</p>
        </div>

        <div className="lg:hidden">
          <Hamburger color="white" toggled={isOpen} size={25} toggle={setIsOpen} />
        </div>
      </div>
      <AnimatePresence>
        {isOpen && (
          <motion.div key="mobile-menu" initial={{ y: "100%" }} animate={{ y: 0 }} exit={{ y: "100%" }} transition={{ type: "spring", stiffness: 100, damping: 20 }}className="fixed inset-0 bg-black w-screen  text-white flex flex-col items-center py-20 z-50">
            <div className="absolute top-5 right-5">
              <Hamburger color="white" toggled={isOpen} toggle={setIsOpen} />
            </div>
            <div className="flex flex-col items-center gap-8 mt-10">
              <Link to="/"><p className="text-3xl font-bold cursor-pointer" onClick={() => setIsOpen(false)}>HOME</p></Link>
              <p className="text-3xl font-bold cursor-pointer" onClick={() => setIsOpen(false)}>PROFILE</p>
              <p className="text-3xl font-bold cursor-pointer" onClick={() => setIsOpen(false)}>ABOUT</p>
            </div>
            <div className='pt-52 flex flex-col gap-y-10 items-center justify-center'>
                <SiGooglegemini size={40}/>
                <p>POWERED BY GEMINI LLM MODEL</p>
            </div>
          </motion.div>
        )}
      </AnimatePresence>
    </>
  );
}
export default NavBar