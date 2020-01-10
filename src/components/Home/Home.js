import React, {useState} from 'react';
import StyledHome from './Home.styled';
import Burger from '../Burger';
import Menu from '../Menu';

const Home = () => {
    const [open, setOpen] = useState(false);
    return (
        <div className="container">
            <div>
                <Burger open={open} setOpen={setOpen}/>
                <Menu open={open} setOpen={setOpen}/>
            </div>

            <div className="row">
                <div className="col-10 text-center pt-5 pl-5">
                    <div >

                        <img className="headshot" src="/static/img/headshot.jpg" alt="cyrus javan headshot"></img>
                    </div>
                    <h1 className="my-3">Cyrus Javan</h1>
                    <h2 className="my-3">Software Engineer</h2>
                    <h2 className="my-3">Currently helping build the world's largest online counseling platform at <a href="https://www.betterhelp.com">BetterHelp</a></h2>
                </div>
                <ul className="navbar-nav mr-auto">
                    <li className="nav-item active">
                        <a className="nav-link social-icon" href="https://github.com/CyrusJavan">
                            <i className="fa fa-github"></i>
                        </a>
                    </li>
                    <li className="nav-item">
                        <a className="nav-link social-icon" href="mailto:javan.cyrus@gmail.com">
                            <i className="fa fa-envelope"></i>

                        </a>
                    </li>
                    <li className="nav-item">
                        <a className="nav-link social-icon" href="https://www.linkedin.com/in/cyrusjavan">
                            <i className="fa fa-linkedin"></i>
                        </a>
                    </li>
                </ul>

            </div>
        </div>
    );
}

export default Home;