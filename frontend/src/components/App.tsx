import React, {useEffect} from "react";
import {BrowserRouter as Router, Route} from "react-router-dom";
import AuthorizationPage from "../pages/AuthorizationPage";
import TasksPage from "../pages/TasksPage";
import MyAppBar from "./MyAppBar";
import ApolloClient from "apollo-client";
import {InMemoryCache} from "apollo-cache-inmemory";
import {ApolloProvider} from "@apollo/react-hooks";
import {createHttpLink} from "apollo-link-http";
import {useDispatch, useSelector} from "react-redux";
import {connectAuth, userSelector} from "../reducers/UserReducer";
import ItemsPage from "../pages/ItemsPage";

const App: React.FC = () => {
    const dispatch = useDispatch();
    useEffect(() => {
        dispatch(connectAuth())
    }, [dispatch]);
    return (
        <ApolloClientProvider>
            <RouteProvider/>
        </ApolloClientProvider>
    );
};

const getApolloClient = (token: string) => {
    const link = createHttpLink({
        uri: 'http://localhost:8080/query',
        headers: {
            authorization: token
        }
    });
    return new ApolloClient({
        link: link,
        cache: new InMemoryCache(),
    });
};

const ApolloClientProvider = (props: any) => {
    const {user} = useSelector(userSelector);
    const token = user && user.token ? `Bearer ${user.token}` : '';
    const client = getApolloClient(token);
    return (
        <ApolloProvider client={client} {...props}/>
    )
};

const RouteProvider = () => {
    return (
        <Router>
            <MyAppBar/>
            <Route path="/" exact component={TasksPage}/>
            <Route path="/users/signin/" exact component={AuthorizationPage}/>
            <Route path="/items/" exact component={ItemsPage}/>
        </Router>
    )
};

export default App;
