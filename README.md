# MultiTenancies
`MultiTenancy` is a kubernetes workload type like `StatefulSet`, except dependant on defined tenants instead of on a replica count.

## Installation

```
kubectl apply -f https://raw.githubusercontent.com/configurator/multitenancy/master/deployment/multitenancy.yaml
```

## Example

And you have a few items deployed:

```yaml
apiVersion: confi.gurator.com/v1
kind: Tenant
tenancyKind: food
metadata:
  name: apple
data:
  name: Apple
  type: Fruit

---

apiVersion: confi.gurator.com/v1
kind: Tenant
tenancyKind: food
metadata:
  name: cucumber
data:
  name: Cucumber
  type: Vegetable

---

apiVersion: confi.gurator.com/v1
kind: Tenant
tenancyKind: food
metadata:
  name: tomato
data:
  name: Tomato
  type: Fruit
```

`MultiTenancies` allow you to create a pod for each of those items. Defining a `MultiTenancy` is much like defining a `Deployment`, or a `StatefulSet`, expect instead of adding a replica count, we tell the operator which tenancyKind to replicate for
```yaml
apiVersion: confi.gurator.com/v1
kind: MultiTenancy
metadata:
  name: food-processor
spec:
  # We want one instace for each Food in our system (required)
  tenancyKind: food
  # Name the env variable for where to put the name of the food (optional)
  tenantNameVariable: FOOD_TYPE
  # Name the volume for where to put the food details (optional)
  tenantResourceVolume: which-food
  # Note: if the tenantResourceVolume is specified, the pod will be restarted for any change in the tenant's data.
  # However, if only tenantNameVariable is specified, the pod will not respond to changes in tenant data

  # Standard selector and template definition, much like for Deployments or StatefulSets:
  selector:
    matchLabels:
      app: food-processor
  template:
    metadata:
      labels:
        app: food-processor
    spec:
      containers:
      - name: food-processor
        image: my-food-processor
        # Add a volume so we can know which Food this pod belongs to
        volumes:
          - name: which-food
            mountPath: /etc/which-food
```

The `MultiTenancy` operator will now create a Pod for each `Tenant` in the system with `tenancyKind: food`, and mount a volume on each one at `/etc/which-food`. Inside that directory, we'll have two files, named `name` and `type`, containing the data from the `Food` definition, much like it would if we had a `ConfigMap`.

If a `Food` item is created or deleted, the pods are killed and started appropriately; for updates, the pod is first killed, then a new one is started.
