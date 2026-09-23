import request from './api';

export const skillService = {
  /**
   * Fetches skills grouped by category from the backend.
   * @param {Object} [options]
   * @param {string} [options.category] Optional category filter
   * @param {boolean} [options.grouped=true] Whether to return grouped categories
   * @returns {Promise<Array>} List of skill categories or raw skills
   */
  async getSkills({ category, grouped = true } = {}) {
    let endpoint = '/skills';
    const params = new URLSearchParams();
    if (category) params.append('category', category);
    if (!grouped) params.append('grouped', 'false');

    const qs = params.toString();
    if (qs) endpoint += `?${qs}`;

    const res = await request(endpoint);
    return res?.data || [];
  },
};

export default skillService;
