import request from './api';

export const projectService = {
  /**
   * Fetches all projects from the Go backend API with optional query parameters.
   * @param {Object} [params]
   * @param {number} [params.page]
   * @param {number} [params.limit]
   * @returns {Promise<Array>} List of projects
   */
  async getProjects(params = {}) {
    const query = new URLSearchParams();
    if (params.page) query.append('page', params.page);
    if (params.limit) query.append('limit', params.limit);
    const qs = query.toString() ? `?${query.toString()}` : '';
    const res = await request(`/projects${qs}`);
    return res?.data || [];
  },

  /**
   * Fetches paginated projects from the Go backend API including pagination metadata.
   * @param {number} page
   * @param {number} limit
   * @returns {Promise<Object>} Full paginated response
   */
  async getPaginatedProjects(page = 1, limit = 3) {
    const res = await request(`/projects?page=${page}&limit=${limit}`);
    return res;
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
