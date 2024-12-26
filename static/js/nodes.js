function nodeManager() {
  return {
      nodes: [],
      loading: false,
      showModal: false,
      isGridView: true,
      newNode: {
          hostname: '',
          user: '',
          publicKey: ''
      },
      isEditing: false,
      editingNodeId: null,
      
      toggleView() {
          this.isGridView = !this.isGridView
          localStorage.setItem('nodesViewPreference', this.isGridView ? 'grid' : 'table')
      },

      init() {
          this.isGridView = localStorage.getItem('nodesViewPreference') !== 'table'
          this.fetchNodes()
      },

      async fetchNodes() {
          this.loading = true
          try {
              const response = await fetch('/nodes')
              const data = await response.json()
              this.nodes = data.nodes
          } catch (error) {
              console.error('Error fetching nodes:', error)
          } finally {
              this.loading = false
          }
      },

      async deleteNode(id) {
          if (!confirm('Are you sure you want to delete this node?')) {
              return
          }
          try {
              await fetch(`/nodes/${id}`, { method: 'DELETE' })
              await this.fetchNodes()
          } catch (error) {
              console.error('Error deleting node:', error)
              alert('Failed to delete node')
          }
      },

      async addNode() {
          try {
              const url = this.isEditing ? `/nodes/${this.editingNodeId}` : '/nodes'
              const method = this.isEditing ? 'PUT' : 'POST'
              
              await fetch(url, {
                  method: method,
                  headers: {
                      'Content-Type': 'application/json',
                  },
                  body: JSON.stringify(this.newNode)
              })
              this.showModal = false
              this.newNode = { hostname: '', user: '', publicKey: '' }
              this.isEditing = false
              this.editingNodeId = null
              await this.fetchNodes()
          } catch (error) {
              console.error('Error saving node:', error)
              alert('Failed to save node')
          }
      },

      editNode(node) {
          this.isEditing = true
          this.editingNodeId = node.ID
          this.newNode = {
              hostname: node.Hostname,
              user: node.User,
              publicKey: node.PublicKey
          }
          this.showModal = true
      }
  }
}