import AppHome from '@/components/AppHome';
const AppMedium = () => import('@/components/AppMedium');
const AppCategories = () => import('@/components/AppCategories');

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
        path: '/categories',
        name: 'categories',
        component: AppCategories
    },
];

export default routes;
