{{ define "Home" }}
<div class="container">
    <div class="row">
        <div class="col-10 text-center pt-5 pl-5">
            <div>
                <img class="headshot" src="/static/img/headshot.jpg" alt="cyrus javan headshot"></img>
            </div>
            <h1 class="my-3">Cyrus Javan</h1>
            <div class="employment">
                <h2>Full Stack Engineer @ <a href="https://www.betterhelp.com">BetterHelp</a></h2>
                <ul>
                    <li>Develop features in PHP and JavaScript for a mental health counseling platform that improve user experience and drive product success.</li>
                    <li>Implemented a quality score for counselors and integrated it into the counselor matching algorithm, prioritizing high quality counselors and increasing revenue per signup by 5%.</li>
                    <li>Built an engagement bar that showed counselors how much they had been interacting with their clients to increase correspondence rates.</li>
                    <li>Identified a malformed header bug causing a production outage, wrote a regular expression to quickly find the source of the issue and deployed a fix to bring back service.</li>
                    <li>Created feature to encourage users to book live counseling sessions, required coordinating multiple cron jobs and user attributes to contact users at appropriate times with the correct messaging.</li>
                    <li>Write and review code on an Agile team, build and test code using a CircleCI pipeline and deploy with Nomad.</li>
                </ul>
            </div>
            <div class="skills">
                <h2>Skills</h2>
                <ul>
                    <li>Languages: PHP, JavaScript, SQL, HTML, CSS, Go, Python, Bash</li>
                    <li>Tools: Git, Linux, AWS (S3, Fargate), Heroku, JetBrains IDEs</li>
                    <li>Frameworks/Libraries: Node.js, ExpressJS, Gin, Docker, Elastic Stack, React, Bootstrap, jQuery</li>
                </ul>
            </div>
            <div class="open-source">
                <h2>Open Source</h2>
                <div>
                    <h3><a href="https://www.github.com/CyrusJavan/portfolio-new">Portfolio Site</a></h3>
                    <ul>
                        <li>Develop in Go using the Gin web framework and deploy on Heroku.</li>
                    </ul>
                </div>
                <div>
                    <h3><a href="https://www.github.com/netauth">NetAuth</a></h3>
                    <ul>
                        <li>Implemented page titles for the GUI of a network level authentication service written in Go.</li>
                        <li>Improved documentation to make it easier to build and run the service.</li>
                    </ul>
                </div>
                <div>
                    <h3><a href="https://mystory.herokuapp.com">MyStory: Online Diary</a></h3>
                    <ul>
                        <li>Built a web app using Node and ExpressJS where users upload stories to be shared publicly or kept
                            private, and comment on others stories, deployed to production with Heroku.</li>
                        <li>Implemented authentication and access control with Passport and Google OAuth2.0.</li>
                        <li>Used MongoDB with MongooseJS to store user information, stories and comments.</li>
                    </ul>
                </div>
            </div>
            <div class="education">
                <h2>Education</h2>
                <h3>BS Computer Science @ San Jose State Univeristy</h3>
            </div>
        </div>


        <ul class="navbar-nav mr-auto">
            <li class="nav-item active">
                <a class="nav-link social-icon" href="https://github.com/CyrusJavan">
                    <i class="fa fa-github"></i>
                </a>
            </li>
            <li class="nav-item">
                <a class="nav-link social-icon" href="mailto:javan.cyrus@gmail.com">
                    <i class="fa fa-envelope"></i>

                </a>
            </li>
            <li class="nav-item">
                <a class="nav-link social-icon" href="https://www.linkedin.com/in/cyrusjavan">
                    <i class="fa fa-linkedin"></i>
                </a>
            </li>
        </ul>
    </div>
</div>
{{end}}