import AppHome from '@/components/AppHome';
import AppStatistic from "./components/AppStatistic";
const AppMedium = () => import('@/components/AppMedium');
const AppCategories = () => import('@/components/AppCategories');
const AppProfile = () => import('@/components/AppProfile');

const routes = [
    {
        path: '/',
        name: 'home',
        component: AppHome
    },
    {
        path: '/medium',
        name: 'Medium',
        component: AppMedium
    },
    {
        path: '/statistic',
        name: 'Statistic',
        component: AppStatistic,
    },
    {
        path: '/categories',
        name: 'categories',
        component: AppCategories
    },
    {
        path: '/profile',
        name: 'profile',
        component: AppProfile
    },
];

export default routes;
