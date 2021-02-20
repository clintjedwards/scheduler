<script>
  import { link, Route, Router } from "svelte-routing";
  import { globalHistory } from "svelte-routing/src/history";
  import AddEmployee from "./components/AddEmployee.svelte";
  import Employees from "./components/Employees.svelte";
  import Footer from "./components/Footer.svelte";
  import ManageEmployee from "./components/ManageEmployee.svelte";
  import Positions from "./components/Positions.svelte";
  import Schedules from "./components/Schedules.svelte";

  let pathname = window.location.pathname;

  globalHistory.listen(({ location, action }) => {
    pathname = location.pathname;
  });

  let routes = [
    { path: "/employees", text: "Employees" },
    { path: "/positions", text: "Positions" },
    { path: "/schedules", text: "Schedules" },
  ];
</script>

<main>
  <div id="content">
    <img id="logo" src="/images/schedule.svg" alt="Scheduler Logo" />
    <h1>Scheduler</h1>
    <Router>
      <nav>
        {#each routes as route}
          <a
            class:selected={route.path === pathname}
            href={route.path}
            use:link
          >
            {route.text}
          </a>
        {/each}
      </nav>
      <div>
        <Route path="employees" component={Employees} />
        <Route path="employees/add" component={AddEmployee} />
        <Route path="employees/:id" component={ManageEmployee} />
        <Route path="positions" component={Positions} />
        <Route path="schedules" component={Schedules} />
      </div>
    </Router>
  </div>
  <div>
    <Footer />
  </div>
</main>

<style>
  :global(html) {
    overflow-y: scroll;
  }

  main {
    display: flex;
    flex-direction: column;
    min-height: 95vh;
  }

  #content {
    width: 65%;
    margin: 0 auto;
    flex: 1;
  }

  main {
    text-align: center;
    padding: 1em;
    max-width: 240px;
    margin: 0 auto;
  }

  #logo {
    max-width: 180px;
  }

  .selected {
    border-bottom: 1px solid #ff3e00;
    color: black;
  }

  h1 {
    color: #ff3e00;
    font-size: 3.5em;
    margin: 0 0 30px 0;
    font-weight: 100;
  }

  nav a {
    font-size: 1.5em;
    text-decoration: none;
    font-weight: 300;
    color: rgb(136 136 136);
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
