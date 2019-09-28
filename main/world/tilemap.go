components {
  id: "tilemap"
  component: "/main/world/floor.tilemap"
  position {
    x: 0.0
    y: 0.0
    z: 0.0
  }
  rotation {
    x: 0.0
    y: 0.0
    z: 0.0
    w: 1.0
  }
}
embedded_components {
  id: "co_bounce"
  type: "collisionobject"
  data: "collision_shape: \"/main/world/floor.tilemap\"\n"
  "type: COLLISION_OBJECT_TYPE_STATIC\n"
  "mass: 0.0\n"
  "friction: 0.5\n"
  "restitution: 3.0\n"
  "group: \"\"\n"
  "mask: \"body\"\n"
  "mask: \"head\"\n"
  "mask: \"jet\"\n"
  "mask: \"chain\"\n"
  "linear_damping: 0.0\n"
  "angular_damping: 0.0\n"
  "locked_rotation: false\n"
  ""
  position {
    x: 0.0
    y: 0.0
    z: 0.0
  }
  rotation {
    x: 0.0
    y: 0.0
    z: 0.0
    w: 1.0
  }
}
embedded_components {
  id: "co_wall"
  type: "collisionobject"
  data: "collision_shape: \"/main/world/floor.tilemap\"\n"
  "type: COLLISION_OBJECT_TYPE_STATIC\n"
  "mass: 0.0\n"
  "friction: 0.5\n"
  "restitution: 0.0\n"
  "group: \"\"\n"
  "mask: \"body\"\n"
  "mask: \"head\"\n"
  "mask: \"jet\"\n"
  "mask: \"chain\"\n"
  "linear_damping: 0.0\n"
  "angular_damping: 0.0\n"
  "locked_rotation: false\n"
  ""
  position {
    x: 0.0
    y: 0.0
    z: 0.0
  }
  rotation {
    x: 0.0
    y: 0.0
    z: 0.0
    w: 1.0
  }
}
