import request from './api';

export const projectService = {
  /**
   * Fetches all projects from the Go backend API.
   * @returns {Promise<Array>} List of projects
   */
  async getProjects() {
    const res = await request('/projects');
    return res?.data || [];
  },

  /**
   * Fetches a single project by its ID.
   * @param {number|string} id
   * @returns {Promise<Object>} Project details
   */
  async getProjectById(id) {
    const res = await request(`/projects/${id}`);
    return res?.data;
  },
};

export default projectService;
