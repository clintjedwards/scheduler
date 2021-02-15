<script>
  import { link, Route, Router } from "svelte-routing";
  import Employees from "./components/Employees.svelte";
  import Footer from "./components/Footer.svelte";
  import Positions from "./components/Positions.svelte";
  import Schedules from "./components/Schedules.svelte";
</script>

<main>
  <div id="mainContent">
    <div id="bodyContent">
      <h1>Scheduler</h1>
      <Router>
        <nav>
          <a href="employees" use:link>Employees</a>
          <a href="positions" use:link>Positions</a>
          <a href="schedules" use:link>Schedules</a>
        </nav>
        <div>
          <Route path="employees" component={Employees} />
          <Route path="positions" component={Positions} />
          <Route path="schedules" component={Schedules} />
        </div>
      </Router>
    </div>
  </div>
  <div>
    <Footer />
  </div>
</main>

<style>
  #mainContent {
    height: 95%;
  }

  #bodyContent {
    width: 65%;
    margin: 0 auto;
  }

  main {
    text-align: center;
    padding: 1em;
    max-width: 240px;
    margin: 0 auto;
  }

  h1 {
    color: #ff3e00;
    font-size: 4em;
    font-weight: 100;
  }

  nav a {
    font-size: 1.5em;
    text-decoration: none;
    font-weight: 100;
    color: black;
    margin-right: 1em;
    margin-left: 1em;
  }

  nav {
    display: flex;
    justify-content: center;
    margin-bottom: 5em;
  }

  @media (min-width: 640px) {
    main {
      max-width: none;
    }
  }
</style>
