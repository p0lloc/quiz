<script>
    import JoinGame from "./view/JoinGame.svelte";
    import QuizList from "./view/QuizList.svelte";
    import {Router, Route} from 'svelte-routing'
</script>

<Router basepath="/app">
    <Route path="/" component={JoinGame} />
    <Route path="/list" component={QuizList} />
</Router>