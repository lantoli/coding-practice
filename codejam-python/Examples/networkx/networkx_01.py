

import networkx as nx
import matplotlib.pyplot as plt

G = nx.Graph()

G.add_nodes_from([1,2,3,4])


G.add_edge(1,2)
G.add_weighted_edges_from([(2,3,4.5)])

print(G.adj)

nx.draw(G)
plt.show()

