import Header from './components/share/Header';
import Footer from './components/share/Footer';
import About from './components/share/About';
import Content from './components/list/Content';


export const routes = [
  { path: '/', component: Content, name:"home"},
  { path:'hakkimizda', component: About , name: "about"}
]
