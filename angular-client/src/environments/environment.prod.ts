export const environment = {
  production: true,
  backend: {
    protocol: '',
    host: '',
    port: '',
    endpoints: {
      allProjectsRequest: '/api/projects',
      oneProjectRequest: '/api/projects/:id'
    }
  }
};
